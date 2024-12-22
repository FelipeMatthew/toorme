package middleware

import (
	"net/http"
	"toorme-api-golang/config"

	"github.com/labstack/echo/v4"
)

func TokenAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token != config.Config.AUTH_TOKEN {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		}
		return next(c)
	}
}
