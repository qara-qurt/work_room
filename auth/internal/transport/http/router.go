package http

import (
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) InitRouter() {

	auth := s.HTTP.Group("/auth")
	{
		auth.POST("/sign-up", s.handler.SignUp)
		auth.POST("/sign-in", s.handler.SignIn)
		auth.GET("/refresh", s.handler.Refresh)
	}

	api := s.HTTP.Group("/api")
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	//check token
	api.Use(s.handler.AuthMiddleware)
	company := api.Group("/company")
	{
		company.POST("", s.handler.CreateCompany)
		company.GET("", s.handler.GetCompanies)
		company.GET("/:id", s.handler.GetCompany)
	}

	users := api.Group("/users")
	{
		users.GET("", s.handler.GetUsers)
		users.GET("/:id", s.handler.GetUser)
	}
}
