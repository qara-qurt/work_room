package http

import (
	"auth/configs"
	"auth/internal/transport/http/handler"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	HTTP    *echo.Echo
	handler *handler.Handler
	config  *configs.Config
}

func NewServer(cfg *configs.Config, handler *handler.Handler) *Server {
	router := echo.New()

	return &Server{
		HTTP:    router,
		handler: handler,
		config:  cfg,
	}
}

func (s *Server) Run() error {
	if err := s.HTTP.Start(fmt.Sprintf(":%s", s.config.Server.Port)); err != nil {
		return err
	}
	logrus.Info(fmt.Sprintf("Server runned on %s", s.config.Server.Port))
	return nil
}
