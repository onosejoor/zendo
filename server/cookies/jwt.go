package cookies

import (
	"fmt"
	"log"
	"main/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SetJwt(payload models.UserRes, exp time.Time) (token string, err error) {

	var secretKey = []byte(os.Getenv("JWT_SECRET"))

	if len(secretKey) < 1 {
		log.Panicln("JWT SECRET CANT BE EMPTY!")
	}

	claim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": payload.Username,
		"id":       payload.ID,
		"exp":      exp.Unix(),
	})

	jwt, err := claim.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return jwt, err
}

func VerifyJwt(session string) (valid bool, err error) {
	var secretKey = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(session, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		return false, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token is valid!")
		fmt.Println("Claims:", claims)
		return true, nil
	} else {
		fmt.Println("Invalid token")
		return false, nil

	}
}

func CreateSession(payload models.UserRes, ctx *fiber.Ctx) error {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	jwt, err := SetJwt(payload, expirationTime)
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Secure:   os.Getenv("ENVIRONMENT") == "production",
		Expires:  expirationTime,
		HTTPOnly: true,
		Name:     "zendo_session_token",
		Value:    jwt,
		SameSite: fiber.CookieSameSiteLaxMode,
		Path:     "/",
	})
	return nil
}
