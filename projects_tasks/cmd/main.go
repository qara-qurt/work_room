package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os/signal"
	"projects_tasks/configs"
	"projects_tasks/internal/repository"
	"projects_tasks/internal/service"
	"projects_tasks/internal/transport/grpc"
	"sync"
	"syscall"
)

func main() {
	log.Fatal(run())
}

func run() error {
	cfg, err := configs.New()
	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("config successfully built")

	repo, err := repository.New(cfg)
	if err != nil {
		log.Error(err)
		return err
	}
	service := service.New(repo, cfg)

	grpcSrv := grpc.NewServer(cfg)
	log.Info("GRPC server has been initialized")

	grpcSrv.RegisterServices(service)
	log.Info("GRPC registered services")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcSrv.Run(); err != nil {
			log.Error(err)
			log.Fatal(err)
		}
	}()
	log.Infof("GRPC server running on %s", cfg.Server.Port)

	gracefulShutdown(grpcSrv, wg)
	wg.Wait()
	return nil
}

func gracefulShutdown(srv *grpc.Server, wg sync.WaitGroup) {
	defer wg.Done()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	srv.GRPCServer.GracefulStop()
	log.Println("Server stopped gracefully")
}
