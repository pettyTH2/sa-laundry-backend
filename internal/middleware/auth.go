package middleware

import (
	"fmt"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"laundry-backend/internal/config"
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
    if err := RequireAuth(c); err != nil {
        return err
    }

    claims := c.Locals("user_claims").(jwt.MapClaims)
    role, _ := claims["role"].(string)

    if role != "cashier" && role != "admin" {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "forbidden: staff only"})
    }

    return c.Next()
}