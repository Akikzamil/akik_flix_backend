package route

import (
	"akikflix/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {
	initPublicRoutes(app)
	middleware.InitJwt(app)
	initPrivateRoutes(app)
}
