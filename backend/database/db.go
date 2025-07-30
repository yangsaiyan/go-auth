package database

import (
	"fmt"
	"go-auth/config"
	"go-auth/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort,
		)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
