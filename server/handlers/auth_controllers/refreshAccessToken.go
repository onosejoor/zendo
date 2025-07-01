package auth_controllers

import (
	"main/cookies"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HandleAccessToken(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("zendo_session_token")
	if refreshToken == "" {
		return ctx.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "unauthorized - no valid access or refresh token",
		})
	}

	newTokens, err := cookies.RefreshAccessToken(refreshToken)
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
		SameSite: "None",
		Path:     "/",
	})

	return ctx.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Access token created",
	})

}
