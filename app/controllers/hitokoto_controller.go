package controllers

import (
	"fmt"
	"slices"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"

	"github.com/TheTNB/hitokoto-api/pkg/hitokoto"
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
	typ, encode, callback, sel, _, _ := sanitize(c)

	sentence, err := hitokoto.Random(typ)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	switch encode {
	case "text":
		return c.SendString(sentence.Hitokoto)
	case "js":
		if callback == "" {
			return c.SendString(fmt.Sprintf(`(function hitokoto(){var hitokoto="%s";var dom=document.querySelector('%s');Array.isArray(dom)?dom[0].innerText=hitokoto:dom.innerText=hitokoto;})()`, sentence.Hitokoto, sel))
		}

		return c.SendString(fmt.Sprintf(`;%s("(function hitokoto(){var hitokoto="%s";var dom=document.querySelector('%s');Array.isArray(dom)?dom[0].innerText=hitokoto:dom.innerText=hitokoto;})()");`, callback, sentence.Hitokoto, sel))
	default:
		if callback != "" {
			encoded, err := sonic.MarshalString(sentence)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": true,
					"msg":   err.Error(),
				})
			}
			json, err := sonic.MarshalString(encoded)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": true,
					"msg":   err.Error(),
				})
			}

			return c.SendString(fmt.Sprintf(`;%s(%s);`, callback, json))
		}

		return c.JSON(sentence)
	}
}

func sanitize(c *fiber.Ctx) (typ, encode, callback, sel string, minLength, maxLength int) {
	types := hitokoto.Types()
	typ = c.Query("c")
	if !slices.Contains(types, typ) {
		typ = ""
	}

	encodes := []string{"text", "json", "js"}
	encode = c.Query("encode", "json")
	if !slices.Contains(encodes, encode) {
		encode = "json"
	}

	callback = c.Query("callback")
	sel = c.Query("select", ".hitokoto")
	minLength = c.QueryInt("min_length", 0)
	maxLength = c.QueryInt("max_length", 0)

	return typ, encode, callback, sel, minLength, maxLength
}
