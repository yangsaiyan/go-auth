package controllers

import (
	"go-auth/services"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx) error {
	email := c.Params("email")
	password := c.Params("password")

	user, err := services.Signup(email, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	email := c.Params("email")
	password := c.Params("password")

	user, err := services.Login(email, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User logged in successfully",
		"user":    user,
	})
}

func Update(c *fiber.Ctx) error {
	email := c.Params("email")
	password := c.Params("password")

	user, err := services.Update(email, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}

func Delete(c *fiber.Ctx) error {
	email := c.Params("email")

	user, err := services.Delete(email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
		"user":    user,
	})
}