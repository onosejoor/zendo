package models

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Email     string             `bson:"email" json:"email"`
	Username  string             `bson:"username" json:"username"`
	Avatar    string             `bson:"avatar" json:"avatar"`
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

func (u *User) ComparePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}
	return true
}

func (p UserPayload) CreateUser(collection *mongo.Collection, ctx context.Context) (id primitive.ObjectID, err error) {
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

type GetUserPayload struct {
	Email    string `bson:"email" json:"email"`
	Username string `bson:"username" json:"username"`
}

func GetUser(userId primitive.ObjectID, collection *mongo.Collection, ctx context.Context) (data GetUserPayload, err error) {
	var user GetUserPayload
	projection := bson.M{"email": 1, "username": 1, "_id": 0}

	opts := options.FindOne().SetProjection(projection)

	err = collection.FindOne(ctx, bson.M{"_id": userId}, opts).Decode(&user)
	if err != nil {
		return GetUserPayload{}, err
	}

	return user, nil
}
