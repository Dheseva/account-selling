package main

import (
	"account-selling/config"
	"account-selling/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {

	engine := html.New("./views", ".html")

	LoadEnv()
	config.Connect()

	app := fiber.New(fiber.Config{
        Views: engine,
    })

	app.Static("/public", "./public")
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(getPort())
}

func getPort() string{
	port := os.Getenv("DB_PORT")
	if port == ""{
		return ":8000"
	}

	return ":" + port
}

func LoadEnv(){
	if err := godotenv.Load(); err != nil{
		log.Fatal("Error load .env file")
	}
}