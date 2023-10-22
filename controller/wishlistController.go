package controller

import (
	"account-selling/config"
	"account-selling/internal/http/middleware"
	modelsuser "account-selling/models/user"
	modelwish "account-selling/models/wishlist"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AddWishlist(c *fiber.Ctx) error {

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

	wish := modelwish.Wishlist{
		Items_id: idInt,
		User_id: int(user.Id),
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&wish)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success register data",
		"data":  wish,
	})
}

func RemoveWishlist(c *fiber.Ctx) error {

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

	var wish modelwish.Wishlist

	wish = modelwish.Wishlist{
		Id: uint(idInt),
		Items_id: wish.Items_id,
		User_id: int(user.Id),
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&wish)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success remove wishlist",
		"data":  wish,
	})
}