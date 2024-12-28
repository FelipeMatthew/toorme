package handler

import (
	"net/http"
	"time"
	"toorme-api-golang/config"
	"toorme-api-golang/internal/domain/models"
	"toorme-api-golang/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte(config.Config.JWT_TOKEN)

type userLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var login userLogin
	if err := c.Bind(&login); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("invalid data"))
	}

	var user models.User
	result := config.DB.Where("username = ?", login.Username).First(&user)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("invalid username"))
		}
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("internal error"))
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.ErrorResponse("invalid password"))
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("error generating JWT token"))
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": signedToken,
		"role":  user.Role,
	})
}
