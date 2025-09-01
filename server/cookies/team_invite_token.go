package cookies

import (
	"fmt"
	"main/configs"
	"main/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	inviteTokenExpiry = 7 * 24 * time.Hour
	inviteSecret      = []byte(os.Getenv("EMAIL_SECRET"))
)

type InviteTokenPayload struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type InviteTokenProps struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

func GenerateTeamInviteEmailToken(payload *models.TeamMemberSchema) (string, error) {
	inviteEmailClaims := &InviteTokenPayload{
		TeamID: payload.TeamID.Hex(),
		UserID: payload.UserID.Hex(),
		Role:   payload.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, inviteEmailClaims)
	return token.SignedString(emailSecret)
}

func VerifyTeamEmailInviteToken(tokenString string) (*InviteTokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &InviteTokenPayload{}, func(token *jwt.Token) (any, error) {
		return emailSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*InviteTokenPayload); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateMagicTeamInviteLink(props EmailProps) string {
	magicLink := fmt.Sprintf("%s/teams/create_team_member?token=%s", os.Getenv("FRONTEND_URL"), props.Token)

	return fmt.Sprintf(configs.TEAM_INVITATION_EMAIL_TEMPLATE, props.Username, magicLink)
}
