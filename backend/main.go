package main

import (
	"github.com/gofiber/fiber/v2"
	"go-auth/config"
	"go-auth/database"
	"go-auth/routes"
	"log"
)

func main() {

	config.LoadConfig()

	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
