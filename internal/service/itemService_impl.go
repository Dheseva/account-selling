package service

import (
	"account-selling/config"
	"account-selling/helper"
	"account-selling/internal/entity"
	"account-selling/internal/http/middleware"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type ItemService struct {
	// Implementasi sesuai kebutuhan Anda
	DB ItemServices
}

func NewItemService(db ItemServices) *ItemService {
    return &ItemService{DB: db}
}

func (db *ItemService) Create(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	Secretkey := helper.PrivateKey()

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

	var user entity.User
	config.DB.Where("id = ?", claims.Issuer).First(&user)

	stockitem, _ := strconv.Atoi(data["stock"])
	itemdata := entity.ItemData{
		Type: data["type"],
		Stock: stockitem,
		Desc: data["desc"],
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&itemdata)

	priceitem, _ := strconv.ParseInt(data["price"], 10, 64)
	item := entity.Items{
		Name: data["name"],
		Price: priceitem,
		Itemdata_id: int(itemdata.Id),
		User_id: int(user.Id),
		Created_at: time.Now().UnixMilli(),
		Updated_at: time.Now().UnixMilli(),
	}
	config.DB.Create(&item)
	
	return config.DB.Create(data).Error
}

// func (db *ItemService) Where(query interface{}, args ...interface{}) ItemServices {
	
// 	return db // Implementasi sesuai kebutuhan Anda
// }

// func (db *ItemService) First(out interface{}, where ...interface{}) ItemServices {
// 	return db // Implementasi sesuai kebutuhan Anda
// }