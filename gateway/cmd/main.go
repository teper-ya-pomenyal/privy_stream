package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/teper-ya-pomenyal/privy_stream/gateway/internal/clients"
	"github.com/teper-ya-pomenyal/privy_stream/gateway/internal/config"
	"github.com/teper-ya-pomenyal/privy_stream/gateway/internal/handlers"
	"github.com/teper-ya-pomenyal/privy_stream/gateway/internal/middlewares"
	"github.com/teper-ya-pomenyal/privy_stream/jwtmanager"
)

func main() {
	cfg := config.LoadConfig()
	userClient, err := clients.NewUserClient(cfg.UserServiceAddress)
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := jwtmanager.LoadPublicKey(cfg.PubKeyAddress)
	if err != nil {
		log.Fatal(err)
	}
	verifier := jwtmanager.NewVerifier(publicKey)
	mw := middlewares.NewMiddleWares(verifier)

	userHandler := handlers.NewUserHandler(userClient)

	userRouter := userHandler.NewRouter(mw)

	srv := http.Server{
		Addr:    ":" + cfg.Port,
		Handler: userRouter,
	}

	go func() {
		log.Printf("gateway listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cansel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cansel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("gateway остановлен по таймауту: %v", err)
	} else {
		log.Print("gateway остановлен gracefully")
	}
}
