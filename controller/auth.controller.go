package controller

import (
	"akikflix/database"
	"akikflix/model"
	"akikflix/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	user2, isUserExist := database.CheckIfUserExists(user.Phone)

	if !isUserExist {
		return c.Status(503).SendString("User does not exist with phone number")
	}

	isPasswordMatched := util.CheckPasswordHash(user.Password, user2.Password)

	if !isPasswordMatched {
		return c.Status(505).SendString("password mismatched!")
	}

	err, token := getToken(user2.Name, user2.IsAdmin)

	if err != nil {
		return c.Status(506).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"token": token})
}

func SendSMS(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	_, isUserExist := database.CheckIfUserExists(user.Phone)

	if isUserExist {
		return c.Status(504).SendString("User already exist with the phone number")
	}

	_, err := twilioSendOTP(user.Phone)
	if err != nil {
		return c.Status(505).SendString(err.Error())
	}

	return c.Status(200).SendString("OTP sent successfully")
}

type OTP struct {
	OTP  string     `json:"otp"`
	User model.User `json:"user"`
}

func SignUp(c *fiber.Ctx) error {
	otp := new(OTP)

	if err := c.BodyParser(otp); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	err := twilioVerifyOTP(otp.User.Phone, otp.OTP)

	if err != nil {
		return c.Status(504).SendString(err.Error())
	}

	user := otp.User

	encryptedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return c.Status(505).SendString(err.Error())
	}

	user.Password = encryptedPassword

	database.CreateUser(&user)

	err, token := getToken(user.Name, user.IsAdmin)

	if err != nil {
		return c.Status(506).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"token": token})
}

func getToken(userName string, isAdmin bool) (error, string) {
	claims := jwt.MapClaims{
		"name":  userName,
		"admin": isAdmin,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);
	secret := util.GetVariable("jwtKey");
	t, err := token.SignedString([]byte(secret));
	if err != nil {
		return err, "";
	}
	return nil, t
}
