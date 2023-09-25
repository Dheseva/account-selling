package controller

import (
	"account-selling/config"
	"account-selling/middleware"
	modelsitem "account-selling/models/items"
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

	stockitem, _ := strconv.Atoi(data["stock"])
	itemdata := modelsitem.ItemData{
		Type:       data["type"],
		Stock:      stockitem,
		Desc:       data["desc"],
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&itemdata)

	priceitem, _ := strconv.ParseInt(data["price"], 10, 64)
	item := modelsitem.Items{
		Name:        data["name"],
		Price:       priceitem,
		Itemdata_id: int(itemdata.Id),
		User_id:     int(user.Id),
		Created_at:  time.Now().UnixMilli(),
		Updated_at:  time.Now().UnixMilli(),
	}
	config.DB.Create(&item)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success register data",
		"data": fiber.Map{
			"item":      item,
			"item_data": itemdata,
		},
	})
}