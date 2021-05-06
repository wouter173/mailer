package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/wouter173/mailer/handlers"
	"github.com/wouter173/mailer/middleware"
	"github.com/wouter173/mailer/utils"
)

func main() {
	godotenv.Load()

	utils.Readkeys()

	app := fiber.New()

	//Disable all cors domains so no one will try to access the api from a frontend.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "",
	}))
	app.Use(middleware.AuthHandler)

	app.All("/send", handlers.SendHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
