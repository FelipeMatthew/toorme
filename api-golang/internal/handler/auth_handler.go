package handler

import (
	"net/http"
	"toorme-api-golang/internal/service"
	"toorme-api-golang/pkg/utils"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

type userLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var login userLogin
	if err := c.Bind(&login); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("invalid data"))
	}

	token, role, err := h.AuthService.Authenticate(login.Username, login.Password)
	if err != nil {
		if err.Error() == "invalid username" || err.Error() == "invalid password" {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("internal error"))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
		"role":  role,
	})
}
