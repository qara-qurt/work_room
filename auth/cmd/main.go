package main

import (
	"auth/configs"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/transport/http"
	"auth/internal/transport/http/handler"
	"log"
)

func main() {
	log.Fatal(run())
}

func run() error {
	config, err := configs.New()
	if err != nil {
		return err
	}

	repo, err := repository.New(config)
	if err != nil {
		return err
	}
	service := service.New(repo)
	handler := handler.New(service)

	srv := http.NewServer(config, handler)
	srv.InitRouter()

	if err := srv.Run(); err != nil {
		return err
	}
	return nil
}
