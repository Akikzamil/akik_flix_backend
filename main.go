package main

import (
	"akikflix/database"
	"akikflix/middleware"
	"akikflix/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDatabase()
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024, // this is the default limit of 4MB
	})
	app.Static("/files", "./files")
	middleware.InitMiddleware(app)
	route.InitRoutes(app)
	app.Listen(":8000")
}
