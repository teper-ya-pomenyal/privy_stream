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
		return nil, fmt.Errorf("неудалось подключиться к Redis: %w", err)
	}
	log.Println(pong)
	return &RedisSessionStore{
		ttlRefresh: ttlRefresh,
		conn:       rdb,
	}, nil
}

func (r *RedisSessionStore) Save(ctx context.Context, refreshToken string, userUUID uuid.UUID) error {
	err := r.conn.Set(ctx, refreshToken, userUUID.String(), r.ttlRefresh).Err()
	if err != nil {
		return err
	}
	return nil
}

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

func (r *RedisSessionStore) Delete(ctx context.Context, refreshToken string) error {
	err := r.conn.Del(ctx, refreshToken).Err()
	if err != nil {
		return err
	}
	return nil
}
