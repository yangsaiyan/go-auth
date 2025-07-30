package utils

import (
	"go-auth/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWTCookie(userID uint) (*fiber.Cookie, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return nil, err
	}

	cookie := &fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(72 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
	}

	return cookie, nil
}

func ClearJWTCookie() *fiber.Cookie {
	return &fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
	}
}
