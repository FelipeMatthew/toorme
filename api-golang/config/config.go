package config

import (
	"log/slog"
	"os"
	"toorme-api-golang/internal/domain/models"
	"toorme-api-golang/pkg/logger"

	"github.com/joho/godotenv"
)

var Config *models.ConfigStruct

func LoadEnv() {
	slog.Info("Loading environment variables")
	err := godotenv.Load(".env")
	logger.LogOnError(err, "Error loading .env file")

	Config = &models.ConfigStruct{
		CTN_VERSION:         os.Getenv("CTN_VERSION"),
		CTN_BUILD:           os.Getenv("CTN_BUILD"),
		DB_USER:             os.Getenv("DB_USER"),
		DB_PASS:             os.Getenv("DB_PASS"),
		DB_HOST:             os.Getenv("DB_HOST"),
		DB_PORT:             os.Getenv("DB_PORT"),
		DB_NAME:             os.Getenv("DB_NAME"),
		AUTH_TOKEN:          os.Getenv("AUTH_TOKEN"),
		CORS_ALLOWED_ORIGIN: os.Getenv("CORS_ALLOWED_ORIGIN"),
	}
}
