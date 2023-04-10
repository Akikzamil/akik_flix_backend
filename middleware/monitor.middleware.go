package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func initMonitor(app *fiber.App) {
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Akikflix Metrics Page"}))
}
