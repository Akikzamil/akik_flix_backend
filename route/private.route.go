package route

import (
	// "akikflix/middleware"
	"akikflix/service"

	"github.com/gofiber/fiber/v2"
)

func initPrivateRoutes(app *fiber.App) {
	restricted := app.Group("/api/v1/private")
	// middleware.InitJwt(restricted)
	app.Get("/restricted", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the AkikFlix!")
	})
	service.FileService(restricted)
}
