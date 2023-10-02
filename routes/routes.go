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
	
	app.Post("/api/complain", controller.AddComplain)
	app.Get("/api/complain/:id", controller.UpdateComplain)
	app.Post("/api/complains/:id", controller.DeleteComplain)

	app.Post("/api/item", controller.AddItems)
	app.Get("/api/item/:id", controller.UpdateItems)
	app.Post("/api/items/:id", controller.DeleteItems)

	app.Get("/api/transaction/:id", controller.AddTransaction)
	app.Get("/api/transactions/:id", controller.UpdateTransaction)
	
}