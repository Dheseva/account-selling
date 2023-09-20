package controller

import (
	"account-selling/config"
	"account-selling/middleware"
	"account-selling/models"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil{
		return err
	}

	var user models.User

	config.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "user not found",
			"data": nil,
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "incorrect password",
			"data": nil,
		})
	}

	var userdata models.UserData
	config.DB.Where("id = ?", user.UData_id).First(&userdata)
	convKey := []byte(Secretkey)

	costumclaims := &middleware.MyCustomClaims{
		IdUser: int(user.Id),
		Name: user.Name,
		StandardClaims: jwt.StandardClaims{
			Issuer: strconv.Itoa(int(user.Id)),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1day
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, costumclaims)

	

	token, err := claims.SignedString(convKey)

	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"status": false,
			"error": err.Error(),
			"message": "could not log in",
			"data": fiber.Map{
				"user": user,
				"user_data": userdata,
			},
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success log in",
		"data": fiber.Map{
			"user": user,
			"user_data": userdata,
		},
	})
}

func User (c *fiber.Ctx) error{

	cookie := c.Cookies("jwt")
 	standClaims := &middleware.MyCustomClaims{}
	convKey := []byte(Secretkey)
	token, err := jwt.ParseWithClaims(cookie, standClaims, func(t *jwt.Token) (interface{}, error) {
		return convKey, nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status": false,
			"error": err.Error(),
			"message": "unauthenticated user",
		})
	}

	claims := token.Claims.(*middleware.MyCustomClaims)

	var user models.User
	config.DB.Where("id = ?",claims.Issuer).First(&user)

	var userdata models.UserData
	config.DB.Where("id = ?", user.UData_id).First(&userdata)

	return c.JSON(fiber.Map{
		"status": true,
		"data": fiber.Map{
			"user": user,
			"user_data": userdata,
		},
	})
}

func Logout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success log out",
	})
}