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
	log.Println("Configure your Extension with the following details:")
	log.Printf("Endpoint URL: http://localhost:%d\n", config.Config.Port)
	log.Printf("Authorization Key: %s\n", config.Config.AuthorizationKey)
	log.Println("Note: Keep the Authorization Key secure and do not share it publicly. The Authorization Key will change each time the server restarts. Use this to reset the key if needed.")
	log.Fatal(app.Listen(":3000"))
}
