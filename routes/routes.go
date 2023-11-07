package routes

import (
	"account-selling/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Setup(app *fiber.App) {
	app.Use(logger.New())

	apiUsers(app)
	apiItems(app)
	apiTransac(app)
	apiComplain(app)
	apiWishlist(app)

}

func apiUsers(app *fiber.App){
	app.Post("/api/register", controller.Register) // how to fix dis, helpme
	app.Post("/api/login", controller.Login)
	app.Get("/api/user", controller.User)
	app.Get("/api/user/:id", controller.UpdateDataUser)
	app.Post("/api/user/topup", controller.TopupUser)
	app.Post("/api/user/delete", controller.DeleteDataUser)
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
	app.Get("/api/transaction/:id", controller.ShowTransaction)
}

func apiComplain(app *fiber.App){
	app.Post("/api/complain", controller.AddComplain)
	app.Get("/api/complain/:id", controller.UpdateComplain)
	app.Post("/api/complains/:id", controller.DeleteComplain)
	app.Post("/api/complains/:id", controller.ShowWishlist)
}

func apiWishlist(app *fiber.App){
	app.Post("/api/wishlist", controller.ShowWishlist)
}