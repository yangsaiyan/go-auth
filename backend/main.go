package main

import (
	"go-auth/config"
	"go-auth/database"
	"go-auth/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	config.LoadConfig()

	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",	
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Cookie, Set-Cookie",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
	}))

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
