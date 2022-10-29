package main

import (
	"beep-work-backend/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	app := fiber.New()

	// Middlewares
	app.Use(logger.New())

	// Routes
	route.InitRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
