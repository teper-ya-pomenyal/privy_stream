package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	userv1 "github.com/teper-ya-pomenyal/privy_stream/proto/user/v1"
	handler "github.com/teper-ya-pomenyal/privy_stream/user_service/internal/delivery/grpc"
	grpc "google.golang.org/grpc"

	"github.com/teper-ya-pomenyal/privy_stream/jwtmanager"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/config"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/repository/postgres"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/repository/redis"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/usecase"
)

func main() {

	cacheDBConfig := config.NewUserDBCacheConfig()
	cfg := config.LoadConfig()
	userDBConfig := config.NewUserDBConfig()

	userRepo, err := postgres.NewUsersPostgresRepository(userDBConfig)
	if err != nil {
		log.Fatal(err)
	}

	userCache, err := redis.NewRedisSessionStore(cacheDBConfig, time.Duration(time.Second*5))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := jwtmanager.LoadPrivateKey("keys/private.pem")
	if err != nil {
		log.Fatal(err)
	}

	tokenManager := jwtmanager.NewManager(privateKey, cfg.TTLAccess)

	login := usecase.NewLoginUseCase(userRepo, tokenManager, userCache)
	register := usecase.NewRegisterUseCase(userRepo, tokenManager, userCache)
	refresh := usecase.NewRefreshUseCase(tokenManager, userCache)
	logout := usecase.NewLogoutUseCase(userCache)

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()

	userHandler := handler.NewUserGRPCHandler(login, register, refresh, logout)

	userv1.RegisterUserServiceServer(grpcServer, userHandler)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	grpcDone := make(chan struct{})

	go func() {
		grpcServer.GracefulStop()
		close(grpcDone)
	}()

	select {
	case <-grpcDone:
		log.Println("gRPC сервер остановлен gracefully")
	case <-time.After(10 * time.Second):
		log.Println("gRPC сервер остановлен по таймауту")
	}

}
