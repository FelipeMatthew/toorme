package middleware

import (
	"net/http"
	"toorme-api-golang/pkg/utils"

	"github.com/labstack/echo/v4"
)

func RoleMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole := c.Get("role").(string)
			for _, role := range allowedRoles {
				if userRole == role {
					return next(c)
				}
			}
			return c.JSON(http.StatusForbidden, utils.ResponseText("error", "denied access, user are not allowed to access"))
		}

	}
}
