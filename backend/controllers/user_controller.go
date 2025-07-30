package controllers

import (
	"fmt"
	"go-auth/services"
	"go-auth/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Signup(c *fiber.Ctx) error {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user, err := services.Signup(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	cookie, err := utils.GenerateJWTCookie(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error generating cookie",
		})
	}
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user, err := services.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	fmt.Printf("Login - User ID: %d\n", user.ID)
	cookie, err := utils.GenerateJWTCookie(user.ID)
	if err != nil {
		fmt.Printf("Login - Error generating cookie: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error generating cookie",
		})
	}
	fmt.Printf("Login - Cookie generated: %+v\n", cookie)
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "User logged in successfully",
		"user":    user,
	})
}

func Update(c *fiber.Ctx) error {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user, err := services.Update(body.Email, body.Password)
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
	body := struct {
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user, err := services.Delete(body.Email)
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

func Logout(c *fiber.Ctx) error {

	cookie := utils.ClearJWTCookie()
	c.Cookie(cookie)

	return c.JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

// To check if the user is logged in
func CheckUser(c *fiber.Ctx) error {
	userClaims := c.Locals("user").(jwt.MapClaims)
	fmt.Printf("CheckUser - User claims: %+v\n", userClaims)
	userID := uint(userClaims["user_id"].(float64))
	fmt.Printf("CheckUser - User ID: %d\n", userID)

	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User fetched successfully",
		"user":    user,
	})
}
