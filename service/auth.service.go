package service

import (
	"akikflix/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthService(app fiber.Router) {
	auth:=app.Group("/auth");
	auth.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the AkikFlix!")
	})
	auth.Post("/login", controller.Login);
	auth.Post("/otp", controller.SendSMS);
	auth.Post("/signUp", controller.SignUp);
}