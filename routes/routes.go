package routes

import (
	handler "account-selling/internal/http/handler"
	"account-selling/mvc/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())

	apiUsers(app)
	apiItems(app)
	apiTransac(app)
	apiComplain(app)

}

func apiUsers(app *fiber.App){
	app.Post("/api/register",handler.RegisterUser)
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Get("/api/user/:id", controller.UpdateDataUser)
	app.Post("/api/logout", controller.Logout)
}

func apiItems(app *fiber.App){
	app.Post("/api/item", controller.AddItems)
	app.Get("/api/item/:id", controller.UpdateItems)
	app.Post("/api/items/:id", controller.DeleteItems)
}

func apiTransac(app *fiber.App){
	app.Get("/api/transaction/:id", controller.AddTransaction)
	app.Get("/api/transactions/:id", controller.UpdateTransaction)
}

func apiComplain(app *fiber.App){
	app.Post("/api/complain", controller.AddComplain)
	app.Get("/api/complain/:id", controller.UpdateComplain)
	app.Post("/api/complains/:id", controller.DeleteComplain)
}