package routes

import (
	"account-selling/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())

	app.Post("/api/register",controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Get("/api/user/:id", controller.UpdateDataUser)
	app.Post("/api/logout", controller.Logout)	
}