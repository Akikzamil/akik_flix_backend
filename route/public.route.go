package route

import (
	"akikflix/service"

	"github.com/gofiber/fiber/v2"
)

func initPublicRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the AkikFlix!")
	})
	version1 := app.Group("/api/v1");
	service.AuthService(version1)

}