package cookies

import (
	"main/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateSession(payload models.UserRes, ctx *fiber.Ctx) error {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	jwts, err := GenerateTokens(payload)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Secure:   os.Getenv("ENVIRONMENT") == "production",
		Expires:  expirationTime,
		HTTPOnly: true,
		Name:     "zendo_session_token",
		Value:    jwts.RefreshToken,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	})

	ctx.Cookie(&fiber.Cookie{
		Secure:   os.Getenv("ENVIRONMENT") == "production",
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Name:     "zendo_access_token",
		Value:    jwts.AccessToken,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	})
	return nil
}
