package main

import (
	"akikflix/database"
	"github.com/gofiber/fiber/v2"
	"akikflix/route"
)

func main() {
	database.InitDatabase()
	app := fiber.New()
	route.InitRoutes(app)
	app.Listen(":8000")
}
