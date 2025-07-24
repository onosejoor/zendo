package cookies

import (
	"errors"
	"fmt"
	"main/configs"
	"main/configs/cron"
	"main/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	tokenExpiry = 15 * time.Minute
	emailSecret = []byte(os.Getenv("EMAIL_SECRET"))
)

type EmailTokenPayload struct {
	ID         string `json:"id"`
	UserName   string `json:"username"`
	EmailToken string `json:"email_token,omitempty"`
	jwt.RegisteredClaims
}

type EmailProps struct {
	Username string
	Token    string
}

func (t *EmailTokenPayload) IsExpired() bool {
	return t.ExpiresAt != nil && t.ExpiresAt.Time.Before(time.Now())
}

func GenerateEmailToken(payload *models.UserRes) (string, error) {
	emailClaims := &EmailTokenPayload{
		ID:       payload.ID.Hex(),
		UserName: payload.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, emailClaims)
	return token.SignedString(emailSecret)
}

func VerifyEmailToken(tokenString string) (*EmailTokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &EmailTokenPayload{}, func(token *jwt.Token) (any, error) {
		return emailSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*EmailTokenPayload); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateMagicLinkEmail(props EmailProps) string {
	magicLink := fmt.Sprintf("%s/auth/verify_email?token=%s", os.Getenv("FRONTEND_URL"), props.Token)

	return fmt.Sprintf(configs.EMAIL_TEMPLATE, props.Username, magicLink)
}

func SendEmail(token string, user models.UserPayload) error {

	template := GenerateMagicLinkEmail(EmailProps{
		Token:    token,
		Username: user.Username,
	})

	err := cron.SendEmailToGmail(user.Email, "Verify your Zendo account", template)
	if err != nil {
		return err
	}
	return nil

}
