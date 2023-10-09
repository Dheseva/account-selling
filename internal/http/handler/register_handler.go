package handler

import (
	"account-selling/internal/entity"
	"account-selling/internal/service"

	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {
	RegisterUserUC *service.RegisterUserUseCase
	RegisterUserdataUC *service.RegisterUserDataUsecase
}

func NewRegisterHandler(registerUUC *service.RegisterUserUseCase, registerUDUC *service.RegisterUserDataUsecase) *RegisterHandler {
	return &RegisterHandler{
		RegisterUserUC: registerUUC,
		RegisterUserdataUC: registerUDUC,
	}
}

func (h *RegisterHandler) RegisterUser(c *fiber.Ctx) error {
	user := new(entity.User)
	userdata := new(entity.UserData)
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	 userdata,err := h.RegisterUserdataUC.Execute(userdata,data); 
	 if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Registration failed",
		})
	}
	// Panggil Use Case untuk registrasi
	if err := h.RegisterUserUC.Execute(user,userdata,data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Registration failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Registration successful",
	})
}