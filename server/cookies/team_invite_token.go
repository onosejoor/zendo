package cookies

import (
	"errors"
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
	models.TeamMemberSchema
	jwt.RegisteredClaims
}

type InviteTokenProps struct {
	Token    string `json:"token"`
	Role     string `json:"role"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
}

func GenerateTeamInviteEmailToken(payload *models.TeamMemberSchema) (string, error) {
	inviteEmailClaims := &InviteTokenPayload{
		TeamMemberSchema: models.TeamMemberSchema{
			TeamID: payload.TeamID,
			UserID: payload.UserID,
			Email:  payload.Email,
			Role:   payload.Role,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(inviteTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, inviteEmailClaims)
	return token.SignedString(inviteSecret)
}

func VerifyTeamEmailInviteToken(tokenString string) (*InviteTokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &InviteTokenPayload{}, func(token *jwt.Token) (any, error) {
		return inviteSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*InviteTokenPayload); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateMagicTeamInviteLink(props InviteTokenProps) string {
	magicLink := fmt.Sprintf("%s/dashboard/teams/members?token=%s", os.Getenv("FRONTEND_URL"), props.Token)

	return fmt.Sprintf(configs.TEAM_INVITATION_EMAIL_TEMPLATE, props.Username, props.TeamName, props.Role, props.Username, magicLink)
}
