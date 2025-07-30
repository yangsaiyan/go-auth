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
    auth.Post("/logout", controllers.Logout)

    user := api.Group("/user", middleware.JwtProtected())
    user.Get("/check", controllers.CheckUser)
}
