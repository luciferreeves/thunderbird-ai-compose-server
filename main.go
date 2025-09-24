package main

import (
	"log"
	"thunderbird-ai-compose-server/config"
	"thunderbird-ai-compose-server/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	routers.Setup(app)

	log.Printf("Starting server on port %d\n", config.Config.Port)
	log.Fatal(app.Listen(":3000"))
}
