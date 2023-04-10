package route

import (
	"akikflix/middleware"
	"github.com/gofiber/fiber/v2"
)

func initPrivateRoutes(app *fiber.App) {
	restricted := app.Group("/api/v1/public")
	middleware.InitJwt(restricted)
	app.Get("/restricted", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the AkikFlix!")
	})

}
