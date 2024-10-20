package main

import (
	"tuj-auth-server/handlers"

	"github.com/labstack/echo/v4"
)

func route(e *echo.Echo) {
	api := e.Group("/api/v1")

	// No JWT auth required
	api.POST("/sign-up", handlers.SignUp)
	api.POST("/login", handlers.Login)

	// JWT auth required
	api.POST("/logout", handlers.Logout)
	api.GET("/validate-token", handlers.ValidateToken)
}
