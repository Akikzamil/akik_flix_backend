package controller

import (
	"akikflix/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	
}

func SignUp(c *fiber.Ctx)error {

}

func getToken(userName string, isAdmin bool) (error, string) {
	claims := jwt.MapClaims{
		"name":  userName,
		"admin": isAdmin,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	secret := util.GetVariable("jwtKey")

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err, ""
	}
	return nil, t
}


