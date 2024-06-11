package controllers

import (
	"github.com/TheTNB/hitokoto-api/pkg/hitokoto"
	"github.com/gofiber/fiber/v2"
)

// Get func get a sentence.
//
//	@Description	Get a sentence.
//	@Summary		get a sentence
//	@Tags			Hitokoto
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	hitokoto.Sentence
//	@Router			/ [get]
func Get(c *fiber.Ctx) error {
	t := c.Params("c")

	sentence, err := hitokoto.Random(t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(sentence)
}
