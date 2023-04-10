package route

import (
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	initPublicRoutes(app)
	initPrivateRoutes(app)
}
