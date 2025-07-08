package cookies

import (
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateSession(payload models.UserRes, ctx *fiber.Ctx) error {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	jwts, err := GenerateTokens(payload)
	if err != nil {
		return err
	}

	isLocal := ctx.Get("Origin") == "http://localhost:3000"

	site := fiber.CookieSameSiteNoneMode

	if isLocal {
		site = fiber.CookieSameSiteLaxMode
	}

	ctx.Cookie(&fiber.Cookie{
		Secure:   !isLocal,
		Expires:  expirationTime,
		HTTPOnly: true,
		Name:     "zendo_session_token",
		Value:    jwts.RefreshToken,
		SameSite: site,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
	})

	ctx.Cookie(&fiber.Cookie{
		Secure:   !isLocal,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Name:     "zendo_access_token",
		Value:    jwts.AccessToken,
		SameSite: site,
		Path:     "/",
		MaxAge:   60 * 15,
	})
	return nil
}
