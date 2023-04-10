package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func InitMiddleware(app *fiber.App) {
	initCacheMiddleware(app);
	initCompressMiddleware(app);
	initCors(app);
	initLogger(app);
	initMonitor(app);
}
