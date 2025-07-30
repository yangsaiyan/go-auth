
package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

var (
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     string
    JWTSecret  string
)

func LoadConfig() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on environment variables")
    }

    DBHost = os.Getenv("DB_HOST")
    DBUser = os.Getenv("DB_USER")
    DBPassword = os.Getenv("DB_PASSWORD")
    DBName = os.Getenv("DB_NAME")
    DBPort = os.Getenv("DB_PORT")
    JWTSecret = os.Getenv("JWT_SECRET")

    if JWTSecret == "" {
        log.Fatal("JWT_SECRET is not set in environment variables")
    }
}
