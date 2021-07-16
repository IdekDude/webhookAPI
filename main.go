package main


import (
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group('/api')

	api.Post('/sendwebhook', )
}