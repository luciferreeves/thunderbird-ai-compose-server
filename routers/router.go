package routers

import (
	"errors"
	"thunderbird-ai-compose-server/config"
	"thunderbird-ai-compose-server/generators"
	"thunderbird-ai-compose-server/generators/gemini"
	"thunderbird-ai-compose-server/types"

	"github.com/gofiber/fiber/v2"
)

func Setup(router *fiber.App) {
	router.Post("/generate", func(c *fiber.Ctx) error {
		var payload types.Payload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(types.ErrorResponse{
				Error: "Invalid request payload",
			})
		}

		prompt := generators.BuildPrompt(payload)

		var response string
		var err error

		switch config.Config.Provider {
		case types.Gemini:
			response, err = gemini.GenerateResponse(prompt)
		default:
			err = errors.New("unsupported AI provider")
		}

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(types.ErrorResponse{
				Error: err.Error(),
			})
		}

		return c.JSON(types.SuccessResponse{
			Response: response,
			Payload:  payload,
		})
	})

	// 404 default
	router.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(types.ErrorResponse{
			Error: "Endpoint not found",
		})
	})
}
