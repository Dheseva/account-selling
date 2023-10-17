package controller

import (
	"account-selling/config"
	"account-selling/internal/http/middleware"
	modelsitem "account-selling/models/items"
	modelsuser "account-selling/models/user"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddItems(c *fiber.Ctx) error {
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
		Type: data["type"],
		Stock: stockitem,
		Desc: data["desc"],
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&itemdata)

	priceitem, _ := strconv.ParseInt(data["price"], 10, 64)
	item := modelsitem.Items{
		Name: data["name"],
		Price: priceitem,
		Itemdata_id: int(itemdata.Id),
		User_id: int(user.Id),
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&item)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success register data",
		"data": fiber.Map{
			"item": item,
			"item_data": itemdata,
		},
	})
}

func UpdateItems(c *fiber.Ctx) error {
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

	var items modelsitem.Items
	config.DB.Where("id = ?", idInt).First(&items)

	var itemdatas modelsitem.ItemData
	config.DB.Where("id = ?", items.Itemdata_id).First(&itemdatas)

	stockitem, _ := strconv.Atoi(data["stock"])
	itemdata := modelsitem.ItemData{
		Id: itemdatas.Id,
		Type: data["type"],
		Stock: stockitem,
		Desc: data["desc"],
		Created_at: itemdatas.Created_at,
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&itemdata)

	priceitem, _ := strconv.ParseInt(data["price"], 10, 64)
	item := modelsitem.Items{
		Id: items.Id,
		Name: data["name"],
		Price: priceitem,
		Itemdata_id: items.Itemdata_id,
		User_id: items.User_id,
		Created_at: items.Created_at,
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&item)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success update data",
		"data": fiber.Map{
			"item": item,
			"item_data": itemdata,
		},
	})
}

func DeleteItems(c *fiber.Ctx) error {

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

	var items modelsitem.Items
	config.DB.Where("id = ?", idInt).First(&items)

	var itemdatas modelsitem.ItemData
	config.DB.Where("id = ?", items.Itemdata_id).First(&itemdatas)

	itemdata := modelsitem.ItemData{
		Id: itemdatas.Id,
		Type: itemdatas.Type,
		Stock: itemdatas.Stock,
		Desc: itemdatas.Desc,
		Created_at: itemdatas.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&itemdata)

	item := modelsitem.Items{
		Id: items.Id,
		Name: items.Name,
		Price: items.Price,
		Itemdata_id: items.Itemdata_id,
		User_id: items.User_id,
		Created_at: items.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&item)

	return c.JSON(fiber.Map{
		"status": true,
		"message": "success delete data",
		"data": fiber.Map{
			"item": item,
			"item_data": itemdata,
		},
	})
}