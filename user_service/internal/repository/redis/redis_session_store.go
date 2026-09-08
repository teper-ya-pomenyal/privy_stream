package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/config"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
)

type RedisSessionStore struct {
	ttlRefresh time.Duration
	conn       *redis.Client
}

func NewRedisSessionStore(cacheDB *config.UserDBCacheConfig, ttlRefresh time.Duration) (*RedisSessionStore, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         cacheDB.DBAddress,
		Password:     cacheDB.DBPassword,
		DB:           cacheDB.DBNumber,
		PoolSize:     cacheDB.PoolSize,
		MinIdleConns: cacheDB.MinIdleConns,
		MaxRetries:   cacheDB.MaxRetries,
		DialTimeout:  cacheDB.DialTimeout,
		ReadTimeout:  cacheDB.ReadTimeout,
		WriteTimeout: cacheDB.WriteTimeout,
	})

	ctx := context.Background()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}
	log.Println(pong)
	return &RedisSessionStore{
		ttlRefresh: ttlRefresh,
		conn:       rdb,
	}, nil
}

func (r *RedisSessionStore) Save(ctx context.Context, refreshToken string, userUUID uuid.UUID) error {
	zsetKey := userUUID.String()

	pipe := r.conn.TxPipeline()
	pipe.Set(ctx, refreshToken, userUUID.String(), r.ttlRefresh)
	pipe.ZAdd(ctx, zsetKey, redis.Z{
		Member: refreshToken,
		Score:  float64(time.Now().Unix()),
	})
	pipe.Expire(ctx, zsetKey, r.ttlRefresh)
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return r.evictIfNeeded(ctx, zsetKey)
}

// Check session
func (r *RedisSessionStore) Get(ctx context.Context, refreshToken string) (uuid.UUID, error) {
	val, err := r.conn.Get(ctx, refreshToken).Result()
	if err == redis.Nil {
		return uuid.UUID{}, domain.ErrRefreshTokenNotFound
	} else if err != nil {
		return uuid.UUID{}, err
	}
	userUUID, err := uuid.Parse(val)
	if err != nil {
		return uuid.UUID{}, err
	}
	return userUUID, nil
}

// Delete session
func (r *RedisSessionStore) Delete(ctx context.Context, refreshToken string) error {
	userUUID, err := r.conn.Get(ctx, refreshToken).Result()
	if err == redis.Nil {
		return domain.ErrRefreshTokenNotFound
	} else if err != nil {
		return err
	}

	pipe := r.conn.TxPipeline()
	pipe.Del(ctx, refreshToken)
	pipe.ZRem(ctx, userUUID, refreshToken)
	if _, err = pipe.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// Refresh sessions
func (r *RedisSessionStore) Refresh(ctx context.Context, oldToken, newToken string) error {
	zsetKey, err := r.conn.Get(ctx, oldToken).Result()
	if err == redis.Nil {
		return domain.ErrRefreshTokenNotFound
	} else if err != nil {
		return err
	}

	pipe := r.conn.TxPipeline()
	pipe.Del(ctx, oldToken)
	pipe.ZRem(ctx, zsetKey, oldToken)
	pipe.Set(ctx, newToken, zsetKey, r.ttlRefresh)
	pipe.ZAdd(ctx, zsetKey, redis.Z{
		Member: newToken,
		Score:  float64(time.Now().Unix()),
	})
	pipe.Expire(ctx, zsetKey, r.ttlRefresh)
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return r.evictIfNeeded(ctx, zsetKey)
}

func (r *RedisSessionStore) evictIfNeeded(ctx context.Context, zsetKey string) error {
	const maxSessions = 7

	count, err := r.conn.ZCard(ctx, zsetKey).Result()
	if err != nil {
		return err
	}
	if count <= maxSessions {
		return nil
	}

	oldToken, err := r.conn.ZRange(ctx, zsetKey, 0, 0).Result()
	if err != nil {
		return err
	}
	if len(oldToken) == 0 {
		return nil
	}

	evictPipe := r.conn.TxPipeline()
	evictPipe.Del(ctx, oldToken[0])
	evictPipe.ZRem(ctx, zsetKey, oldToken[0])
	_, err = evictPipe.Exec(ctx)
	return err
}
