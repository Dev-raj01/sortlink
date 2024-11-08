package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Dev-raj01/sortlink/database"
	"github.com/Dev-raj01/sortlink/handlers"
)

func main() {
	app := fiber.New(fiber.Config{})
	database.InitDB()
	handlers.InitRouter(app)
	app.Listen(":5566")
}
