package middlewares

import (
	"main/cookies"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	accessToken := c.Cookies("zendo_access_token")

	if accessToken != "" {
		user, err := cookies.VerifyAccessToken(accessToken)
		if err == nil {
			c.Locals("user", user)
			return c.Next()
		}
	}

	refreshToken := c.Cookies("zendo_session_token")
	if refreshToken == "" {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "unauthorized - no valid access or refresh token",
		})
	}

	newTokens, err := cookies.RefreshAccessToken(refreshToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": true,
			"message": "invalid refresh token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "zendo_access_token",
		Value:    newTokens.AccessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   os.Getenv("ENVIRONMENT") == "production",
		SameSite: "Lax",
		Path:     "/",
	})

	user, _ := cookies.VerifyAccessToken(newTokens.AccessToken)
	c.Locals("user", user)

	return c.Next()
}
