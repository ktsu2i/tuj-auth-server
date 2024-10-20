package main

import (
	"tuj-auth-server/handlers"

	"github.com/labstack/echo/v4"
)

func route(e *echo.Echo) {
	api := e.Group("/api/v1")

	api.GET("/", handlers.Root)
}
