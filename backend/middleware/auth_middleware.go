package middleware

import (
	"go-auth/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Cookies("jwt")

		if tokenString == "" {
			return c.Status(401).JSON(fiber.Map{
				"error": "Authentication required",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user", claims)
		return c.Next()
	}
}
