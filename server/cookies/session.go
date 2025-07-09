package cookies

import (
	"log"
	"main/models"
	"os"
	"strings"
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
	domain := strings.Split(strings.Split(ctx.Get("Origin"), "//")[1], ":")[0]

	site := fiber.CookieSameSiteNoneMode

	if os.Getenv("ENVIRONMENT") != "production" {
		site = fiber.CookieSameSiteLaxMode
	}

	log.Println(domain)

	ctx.Cookie(&fiber.Cookie{
		Secure:   !isLocal,
		Expires:  expirationTime,
		HTTPOnly: true,
		Name:     "zendo_session_token",
		Value:    jwts.RefreshToken,
		SameSite: site,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		Domain:   domain,
	})

	ctx.Cookie(&fiber.Cookie{
		Secure:   !isLocal,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Name:     "zendo_access_token",
		Value:    jwts.AccessToken,
		SameSite: site,
		Path:     "/",
		Domain:   domain,
		MaxAge:   60 * 15,
	})
	return nil
}
