package http

import "github.com/labstack/echo/v4/middleware"

func (s *Server) InitRouter() {
	api := s.HTTP.Group("/api")
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", s.handler.SignUp)
		auth.POST("/sign-in", s.handler.SignIn)
	}

}
