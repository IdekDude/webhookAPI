package main

import (
	"fmt"
	"github.com/IdekDude/webhookAPI/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api")

	api.Post("/sendwebhook", handlers.SendWebhook)

	port := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%v", port))
}