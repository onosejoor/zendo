package auth_controllers

import (
	"log"
	"main/cookies"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AccesTokenPayload struct {
	RefreshToken string `json:"refresh_token"`
}

func HandleAccessToken(ctx *fiber.Ctx) error {
	var body AccesTokenPayload

	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error parsing body,", err)
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	newTokens, err := cookies.RefreshAccessToken(body.RefreshToken)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "invalid refresh token",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "zendo_access_token",
		Value:    newTokens.AccessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   os.Getenv("ENVIRONMENT") == "production",
		SameSite: "Lax",
		Path:     "/",
	})

	return ctx.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Access token created",
	})

}
