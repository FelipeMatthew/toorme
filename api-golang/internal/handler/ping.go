package handler

import (
	"net/http"
	"toorme-api-golang/config"

	"github.com/labstack/echo/v4"
)

func Ping(c echo.Context) error {
	envInfo := map[string]string{
		"CTN_VERSION":          config.Config.CTN_VERSION,
		"CTN_BUILD":            config.Config.CTN_BUILD,
		"DB_USER":              config.Config.DB_USER,
		"DB_PASS":              "*",
		"DB_HOST":              config.Config.DB_HOST,
		"DB_PORT":              config.Config.DB_PORT,
		"DB_NAME":              config.Config.DB_NAME,
		"AUTH_TOKEN":           "*",
		"CORS_ALLOWED_ORIGINS": config.Config.CORS_ALLOWED_ORIGINS,
		"STATUS":               "OK",
	}

	return c.JSON(http.StatusOK, envInfo)
}
