package util

import (
	"akikflix/model"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUser(c *fiber.Ctx) model.User {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	isAdmin := claims["admin"].(bool)
	return model.User{Name: name, IsAdmin: isAdmin}
}