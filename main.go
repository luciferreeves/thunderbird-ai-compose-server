package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/generate", func(c *fiber.Ctx) error {
		var payload Payload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request payload",
			})
		}

		return c.JSON(fiber.Map{
			"response": "Payload received successfully",
			"data":     payload,
		})
	})

	log.Printf("Starting server on port %d\n", Config.Port)
	log.Fatal(app.Listen(":3000"))
}
