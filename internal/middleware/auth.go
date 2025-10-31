package middleware

import (
	"fmt"
	"laundry-backend/internal/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtConfig = config.LoadJWTConfig()
var jwtSecretKey = []byte(jwtConfig.SecretKey)

func RequireAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "กรุณา login"})
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid claims"})
	}

	c.Locals("user_claims", claims)
	return c.Next()
}

func RequireStaffAuth(c *fiber.Ctx) error {
	claims := c.Locals("user_claims")
	if claims == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "กรุณา login"})
	}

	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid claims"})
	}

	role, _ := mapClaims["role"].(string)

	if role != "cashier" && role != "admin" && role != "laundryAttendant" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden: staff only"})
	}

	return c.Next()
}
