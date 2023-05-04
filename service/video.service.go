package service

import (
	"akikflix/controller"

	"github.com/gofiber/fiber/v2"
)

func FileService(app fiber.Router){
	app.Post("/file",controller.UploadVideo)
	app.Get("/file/:id",controller.GetVideo)
	app.Get("/file",controller.GetVideos)
}