package main

import (
	"auth/configs"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/transport/http"
	"auth/internal/transport/http/handler"
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	logrus.Fatal(run())
}

func run() error {
	config, err := configs.New()
	logrus.Info("config successfully built")
	if err != nil {
		return err
	}

	repo, err := repository.New(config)
	if err != nil {
		logrus.Error(err)
		return err
	}
	service := service.New(repo, config)
	handler := handler.New(service, config)

	srv := http.NewServer(config, handler)
	srv.InitRouter()
	logrus.Info("routes successfully added")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := srv.Run(); err != nil {
			logrus.Error(err)
			logrus.Fatal(err)
		}
	}()
	gracefulShutdown(srv, wg)

	wg.Wait()
	return nil
}

func gracefulShutdown(srv *http.Server, wg sync.WaitGroup) {
	defer wg.Done()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.HTTP.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server stopped gracefully")
}
