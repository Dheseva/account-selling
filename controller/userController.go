package controller

import (
	"account-selling/config"
	"account-selling/internal/http/middleware"
	modelsuser "account-selling/models/user"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func UpdateDataUser(c *fiber.Ctx) error {
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

	var userdata modelsuser.UserData
	config.DB.Where("id = ?", user.UData_id).First(&userdata)

	useridInt := int(user.Id)
	if idInt != useridInt {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "unauthenticated user",
		})
	}
	
	var dateInput int64
	if data["dateofbirth"] != "" {
		convertDate := data["dateofbirth"]
		parsedDate, _ := time.Parse("2006-01-02", convertDate)
		dateInput = parsedDate.UnixNano() / int64(time.Millisecond)
	}else{
		dateInput = 0
	}

	userdata = modelsuser.UserData{
		Id:          userdata.Id,
		Nickname:    data["nickname"],
		Firstname: 	data["firstname"],
		Lastname: 	data["lastname"],
		Sex:         data["sex"],
		Address:     data["address"],
		Dateofbirth: dateInput,
		Nationality: data["nationality"],
		Saldo: 	userdata.Saldo,
		Created_at:  userdata.Created_at,
		Updated_at:  time.Now().UnixMilli(),
		Deleted_at:  userdata.Deleted_at,
	}

	config.DB.Save(&userdata)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success update data user",
		"data": fiber.Map{
			"user":      user,
			"user_data": userdata,
		},
	})
}

func TopupUser(c *fiber.Ctx) error {

	var data map[string]int64
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

	var userdata modelsuser.UserData
	config.DB.Where("id = ?", user.UData_id).First(&userdata)

	var total int64
	if userdata.Saldo != 0{
		total = userdata.Saldo + data["saldo"]
	}else{
		total = data["saldo"]
	}

	userdata = modelsuser.UserData{
		Id:          userdata.Id,
		Nickname:    userdata.Nickname,
		Firstname: 	userdata.Firstname,
		Lastname: 	userdata.Lastname,
		Sex:         userdata.Sex,
		Address:    userdata.Address,
		Dateofbirth: userdata.Dateofbirth,
		Nationality: userdata.Nationality,
		Saldo: 	total,
		Created_at:  userdata.Created_at,
		Updated_at:  time.Now().UnixMilli(),
		Deleted_at:  time.Now().UnixMilli(),
	}
	config.DB.Save(&userdata)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success top up data user",
		"data": fiber.Map{
			"user":      user,
			"user_data": userdata,
		},
	})
}

func DeleteDataUser(c *fiber.Ctx) error {

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

	var userdata modelsuser.UserData
	config.DB.Where("id = ?", user.UData_id).First(&userdata)

	userdata = modelsuser.UserData{
		Id:          userdata.Id,
		Nickname:    userdata.Nickname,
		Firstname: 	userdata.Firstname,
		Lastname: 	userdata.Lastname,
		Sex:         userdata.Sex,
		Address:    userdata.Address,
		Dateofbirth: userdata.Dateofbirth,
		Nationality: userdata.Nationality,
		Saldo: 	userdata.Saldo,
		Created_at:  userdata.Created_at,
		Updated_at:  time.Now().UnixMilli(),
		Deleted_at:  time.Now().UnixMilli(),
	}
	config.DB.Save(&userdata)

	user = modelsuser.User{
		Id: user.Id,
		Name: user.Name,
		Password: user.Password,
		Email: user.Email,
		UData_id: user.UData_id,
		Lastlogin: user.Lastlogin,
		Created_at: user.Created_at,
		Updated_at: time.Now().UnixMilli(),
		Deleted_at: time.Now().UnixMilli(),
	}
	config.DB.Save(&user)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success delete data user",
		"data": fiber.Map{
			"user":      user,
			"user_data": userdata,
		},
	})
}

func ShowallUser(c *fiber.Ctx) error {

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

	var alluser []modelsuser.User
	config.DB.Find(&alluser)

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success remove wishlist",
		"data":  alluser,
	})
}