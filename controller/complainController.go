package controller

import (
	"account-selling/config"
	"account-selling/internal/http/middleware"
	modelcom "account-selling/models/complain"
	modelsuser "account-selling/models/user"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddComplain(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "invalid request",
			"data": nil,
		})
	}

	cookie := c.Cookies("jwt")
	standClaims := &middleware.MyCustomClaims{}
	convKey := []byte(Secretkey)
	token, err := jwt.ParseWithClaims(cookie, standClaims, func(t *jwt.Token) (interface{}, error) {
		return convKey, nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  false,
			"error":   err.Error(),
			"message": "unauthenticated user",
		})
	}

	claims := token.Claims.(*middleware.MyCustomClaims)

	var user modelsuser.User
	config.DB.Where("id = ?", claims.Issuer).First(&user)

	complen := modelcom.Complain{
		Items_id: idInt,
		User_id: int(user.Id),
		Complain: data["complain"],
		Status: 1,
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&complen)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success register data",
		"data": complen,
	})
}

func UpdateComplain(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "invalid request",
			"data": nil,
		})
	}

	cookie := c.Cookies("jwt")
	standClaims := &middleware.MyCustomClaims{}
	convKey := []byte(Secretkey)
	token, err := jwt.ParseWithClaims(cookie, standClaims, func(t *jwt.Token) (interface{}, error) {
		return convKey, nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  false,
			"error":   err.Error(),
			"message": "unauthenticated user",
		})
	}

	claims := token.Claims.(*middleware.MyCustomClaims)

	var user modelsuser.User
	config.DB.Where("id = ?", claims.Issuer).First(&user)

	var idcomplain modelcom.Complain
	config.DB.Where("id = ?", idInt).First(&idcomplain)

	statuse , _ := strconv.Atoi(data["status"])

	complen := modelcom.Complain{
		Id: idcomplain.Id,
		Items_id: idInt,
		User_id: int(user.Id),
		Complain: data["complain"],
		Status: statuse,
		Created_at: idcomplain.Created_at,
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&complen)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success register data",
		"data": complen,
	})
}

func DeleteComplain(c *fiber.Ctx) error {

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status": false,
			"message": "invalid request",
			"data": nil,
		})
	}

	cookie := c.Cookies("jwt")
	standClaims := &middleware.MyCustomClaims{}
	convKey := []byte(Secretkey)
	token, err := jwt.ParseWithClaims(cookie, standClaims, func(t *jwt.Token) (interface{}, error) {
		return convKey, nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"status":  false,
			"error":   err.Error(),
			"message": "unauthenticated user",
		})
	}

	claims := token.Claims.(*middleware.MyCustomClaims)

	var user modelsuser.User
	config.DB.Where("id = ?", claims.Issuer).First(&user)

	var idcomplain modelcom.Complain
	config.DB.Where("id = ?", idInt).First(&idcomplain)

	complen := modelcom.Complain{
		Id: idcomplain.Id,
		Items_id: idcomplain.Items_id,
		User_id: idcomplain.User_id,
		Complain: idcomplain.Complain,
		Status: idcomplain.Status,
		Created_at: idcomplain.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&complen)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success register data",
		"data": complen,
	})
}