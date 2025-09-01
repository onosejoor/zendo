package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func ParseBodyAndValidateStruct(v any, ctx *fiber.Ctx) error {
	if err := ctx.BodyParser(&v); err != nil {
		return errors.New("Invalid request body, error: " + err.Error())
	}
	if err := Validate.Struct(v); err != nil {
		return errors.New("Validation failed, error: " + err.Error())
	}
	return nil
}
