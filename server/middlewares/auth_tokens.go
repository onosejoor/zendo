package middlewares

import (
	"main/cookies"

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

		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Invalid access token",
		})
	}

	return c.Status(401).JSON(fiber.Map{
		"success": false,
		"message": "No access token provided",
	})

	// This code was commented because i learned that refresh tokens should only be in the client

	// refreshToken := c.Cookies("zendo_session_token")
	// if refreshToken == "" {
	// 	return c.Status(401).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "unauthorized - no valid access or refresh token",
	// 	})
	// }

	// newTokens, err := cookies.RefreshAccessToken(refreshToken)
	// if err != nil {
	// 	return c.Status(401).JSON(fiber.Map{
	// 		"success": false,
	// 		"message": "invalid refresh token",
	// 	})
	// }

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "zendo_access_token",
	// 	Value:    newTokens.AccessToken,
	// 	Expires:  time.Now().Add(15 * time.Minute),
	// 	HTTPOnly: true,
	// 	Secure:   os.Getenv("ENVIRONMENT") == "production",
	// 	SameSite: "None",
	// 	Path:     "/",
	// })

	// user, _ := cookies.VerifyAccessToken(newTokens.AccessToken)
	// c.Locals("user", user)

	// return c.Next()
}
