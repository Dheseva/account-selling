package controller

import (
	"account-selling/config"
	"account-selling/middleware"
	modelitem "account-selling/models/items"
	modelTrasac "account-selling/models/transaction"
	modelsuser "account-selling/models/user"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddTransaction(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "invalid request",
			"data":    nil,
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

	var item modelitem.Items
	config.DB.Where("id = ?", idInt).First(&item)

	transac := modelTrasac.Transaction{
		Items_id: int(item.Id),
		Selluser_id: item.User_id,
		Buyuser_id: int(user.Id),
		Price: item.Price,
		Comment: data["comment"],
		Status: 1,
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&transac)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success register data",
		"data":    transac,
	})
}

func UpdateTransaction(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	id := c.Params("id")
	idInt, err := strconv.Atoi(id)

	status, _ := strconv.Atoi(data["status"])

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "invalid request",
			"data":    nil,
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

	var transac modelTrasac.Transaction
	config.DB.Where("id = ?", idInt).First(&transac)

	var user modelsuser.User
	config.DB.Where("id = ?", claims.Issuer).First(&user)

	var item modelitem.Items
	config.DB.Where("id = ?", transac.Items_id).First(&item)

	transac = modelTrasac.Transaction{
		Id: transac.Id,
		Items_id: transac.Items_id,
		Selluser_id: transac.Selluser_id,
		Buyuser_id: transac.Buyuser_id,
		Price: transac.Price,
		Comment: data["comment"],
		Status: status,
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&transac)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success update data",
		"data":    transac,
	})
}