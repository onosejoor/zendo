package cookies

import (
	"errors"
	"log"
	"main/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type Claims struct {
	ID       any    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	accessTokenSecret  = []byte(os.Getenv("ACCESS_SECRET"))
	refreshTokenSecret = []byte(os.Getenv("JWT_SECRET"))
	accessTokenExpiry  = 15 * time.Minute
	refreshTokenExpiry = 7 * 24 * time.Hour
)

func GenerateTokens(user models.UserRes) (*TokenPair, error) {
	accessClaims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(accessTokenSecret)
	if err != nil {
		return nil, err
	}

	refreshClaims := &Claims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    int64(accessTokenExpiry.Seconds()),
	}, nil
}

func VerifyAccessToken(tokenString string) (*models.UserRes, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return accessTokenSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userId, _ := primitive.ObjectIDFromHex(claims.ID.(string))
		return &models.UserRes{
			ID:       userId,
			Username: claims.Username,
		}, nil
	}

	return nil, errors.New("invalid token")
}

func verifyRefreshToken(tokenString string) (*models.UserRes, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return refreshTokenSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userId, _ := primitive.ObjectIDFromHex(claims.ID.(string))
		return &models.UserRes{
			ID:       userId,
			Username: claims.Username,
		}, nil
	}

	return nil, errors.New("invalid refresh token")
}

func RefreshAccessToken(refreshTokenString string) (*TokenPair, error) {
	user, err := verifyRefreshToken(refreshTokenString)
	if err != nil {
		log.Println("Error verifying refresh token: ", err.Error())
		return nil, err
	}

	return GenerateTokens(*user)
}
