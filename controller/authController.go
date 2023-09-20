package controller

import (
	"account-selling/config"
	"account-selling/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var Secretkey = os.Getenv("PRIVATE_KEY_JWT")

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	userdata := models.UserData{
		Nickname: data["name"],
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&userdata)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name: data["name"],
		Password: password,
		Email: data["email"],
		UData_id: int(userdata.Id),
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	

	config.DB.Create(&user)
	return c.JSON(fiber.Map{
		"status": true,
		"message": "success register data",
		"data": fiber.Map{
			"user": user,
			"user_data": userdata,
		},
	})
}