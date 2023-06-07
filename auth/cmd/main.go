package main

import (
	"auth/configs"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/transport/http"
	"auth/internal/transport/http/handler"
	"github.com/sirupsen/logrus"
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
	if err := srv.Run(); err != nil {
		logrus.Error()
		return err
	}
	return nil
}
