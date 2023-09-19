package main

import (
	"account-selling/config"
	"account-selling/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	LoadEnv()
	config.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(GetPort())
}

func GetPort() string{
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