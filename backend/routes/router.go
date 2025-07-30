package routes

import (
    "go-auth/controllers"
    "go-auth/middleware"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api")

    auth := api.Group("/auth")
    auth.Post("/signup", controllers.Signup)
    auth.Post("/login", controllers.Login)

    user := api.Group("/user", middleware.JwtProtected())
    user.Put("/", controllers.Update)
    user.Delete("/", controllers.Delete)
}
