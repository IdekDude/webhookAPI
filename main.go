package main

import (
	"fmt"
	"github.com/IdekDude/webhookAPI/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	api.Post("/sendwebhook", handlers.SendWebhook)


	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%v", port))
}