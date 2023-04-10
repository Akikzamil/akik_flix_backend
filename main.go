package main

import (
	"akikflix/database"
	"akikflix/middleware"
	"akikflix/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New()
	middleware.InitMiddleware(app)
	route.InitRoutes(app)
	app.Listen(":8000")
}
