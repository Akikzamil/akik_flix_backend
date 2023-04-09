package route

import "github.com/gofiber/fiber/v2"

func initPrivateRoutes(app *fiber.App) {
	app.Get("/restricted", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the AkikFlix!")
	})

}