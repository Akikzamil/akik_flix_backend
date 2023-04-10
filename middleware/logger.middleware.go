package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func initLogger(app *fiber.App) {
	app.Use(logger.New(logger.ConfigDefault))
}
