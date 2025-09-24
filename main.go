package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Identity struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ComposeDetails struct {
	Subject    string   `json:"subject,omitempty"`
	To         []string `json:"to,omitempty"`
	Cc         []string `json:"cc,omitempty"`
	Bcc        []string `json:"bcc,omitempty"`
	BodyPlain  string   `json:"bodyPlain,omitempty"`
	BodyHTML   string   `json:"bodyHTML,omitempty"`
	IdentityID string   `json:"identityId,omitempty"`
	IsHTML     bool     `json:"isHTML"`
}

type ComposeContext struct {
	Account Identity       `json:"account"`
	Compose ComposeDetails `json:"compose"`
}

type Payload struct {
	Prompt  string         `json:"prompt"`
	Context ComposeContext `json:"context"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET, HEAD, PUT, PATCH, POST, DELETE, OPTIONS",
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization, X-Requested-With, X-API-Key, X-CSRF-Token",
		ExposeHeaders: "Content-Length, Content-Type, Content-Disposition, X-Pagination, X-Total-Count",
		MaxAge:        86400,
	}))

	app.Post("/generate", func(c *fiber.Ctx) error {
		var payload Payload

		// DEBUG: Simulate processing delay
		time.Sleep(5 * time.Second)

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

	log.Println("Server is running on http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
