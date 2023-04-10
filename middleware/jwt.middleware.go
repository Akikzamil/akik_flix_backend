package middleware

import (
	"akikflix/util"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func InitJwt(app fiber.Router){
	secret := util.GetVariable("akikflix")
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),	
	}))
}