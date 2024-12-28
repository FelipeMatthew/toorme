package middleware

import (
	"net/http"
	"strings"
	"toorme-api-golang/config"
	"toorme-api-golang/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// TODO: trasnform to enviroment variable
var SecretKey = []byte(config.Config.JWT_TOKEN)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("missing or invalid token"))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.NewHTTPError(http.StatusUnauthorized, "unexpected signing method")
			}
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("invalid token"))
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("invalid token claims"))
		}

		c.Set("username", claims["username"])
		c.Set("role", claims["role"])

		return next(c)
	}
}
