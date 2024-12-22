package main

import (
	"net/http"
	"toorme-api-golang/config"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnv()

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "pong")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
