package service

import (
	"errors"
	"time"
	"toorme-api-golang/config"
	"toorme-api-golang/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Authenticate(username, password string) (string, string, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		if err.Error() == "record not found" {
			return "", "", errors.New("invalid username")
		}
		return "", "", errors.New("internal error")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Config.JWT_TOKEN))
	if err != nil {
		return "", "", errors.New("error generating JWT token")
	}

	return signedToken, user.Role, nil
}
