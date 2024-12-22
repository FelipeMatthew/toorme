package router

import (
	"toorme-api-golang/internal/app/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/ping", handler.Ping)

	user := e.Group("user")
	user.GET("/ping", handler.Ping)
}
