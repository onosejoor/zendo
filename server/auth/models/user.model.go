package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Email     string             `bson:"email" json:"email"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"-"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type UserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRes struct {
	Username string `json:"username"`
	ID       any    `json:"id"`
}

func CreateUser(p UserPayload, collection *mongo.Collection, ctx context.Context) (id any, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), 12)
	if err != nil {
		log.Println("error hasing password:", err)
		return "", err
	}

	data, err := collection.InsertOne(ctx, User{
		Email:     p.Email,
		Username:  p.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Insert error:", err)
		return "", err
	}

	return data.InsertedID, nil
}
