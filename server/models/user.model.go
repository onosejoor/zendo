package models

import (
	"context"
	"log"
	"strings"
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
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserRes struct {
	Username string             `json:"username"`
	ID       primitive.ObjectID `json:"id"`
}

func CreateUser(p UserPayload, collection *mongo.Collection, ctx context.Context) (id primitive.ObjectID, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), 12)
	if err != nil {
		log.Println("error hasing password:", err)
		return primitive.NilObjectID, err
	}

	data, err := collection.InsertOne(ctx, User{
		Email:     p.Email,
		Username:  strings.TrimSpace(p.Username),
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Insert error:", err)
		return primitive.NilObjectID, err
	}

	return data.InsertedID.(primitive.ObjectID), nil
}
