package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"laundry-backend/internal/config"
)

var jwtConfig = config.LoadJWTConfig()
var jwtSecretKey = []byte(jwtConfig.SecretKey)

func CreateMemberAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
      return []byte(jwtSecretKey), nil
   	})

	if err != nil || !token.Valid {
      return c.SendStatus(fiber.StatusUnauthorized)
  	}

	claim := token.Claims.(jwt.MapClaims)
	if claim["role"] == "member" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}