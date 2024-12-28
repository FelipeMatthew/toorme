package config

import (
	"fmt"
	"log"
	"log/slog"
	"toorme-api-golang/internal/domain/models"
	"toorme-api-golang/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo",
		Config.DB_HOST,
		Config.DB_USER,
		Config.DB_PASS,
		Config.DB_NAME,
		Config.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.LogOnError(err, "Failed to connect to database")
		log.Fatalf("Could not connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.Supplier{},
		&models.Location{},
		&models.TravelPlan{},
		&models.TravelPlanLocation{},
		&models.Trip{},
		&models.Notification{},
	)
	if err != nil {
		logger.LogOnError(err, "error while automigrate")
	}

	DB = db
	slog.Info("Database connection established")
}
