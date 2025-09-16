package auth_controllers

import (
	"errors"
	"log"
	"main/configs"
	prometheus "main/configs/prometheus"
	"main/cookies"
	"main/db"
	"main/models"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleOauth(ctx *fiber.Ctx, body configs.GooglePayload) (bool, int) {
	collection := db.GetClient().Collection("users")

	var userData models.User
	start := time.Now()
	if err := collection.FindOne(ctx.Context(), bson.M{"email": body.Email}).Decode(&userData); err != nil {
		dbDuration := time.Since(start)
		prometheus.RecordDatabaseQueryOperation(dbDuration, "users", "findOne")
		if errors.Is(err, mongo.ErrNoDocuments) {
			err := createUser(body, collection, ctx)
			if err != nil {
				log.Println("error creating user: ", err)
				ctx.Status(500).JSON(fiber.Map{
					"success": false, "message": "Internal error",
				})
				return true, 500
			}

			return false, 201
		}
		log.Printf("Database error when finding user with email %s: %v\n", body.Email, err)
		ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error getting user data",
		})
		return true, 500
	}

	if !userData.EmailVerified {
		updateStart := time.Now()
		if _, err := collection.UpdateOne(ctx.Context(), bson.M{"_id": userData.ID}, bson.M{
			"$set": bson.M{
				"avatar":        body.Picture,
				"emailVerified": body.EmailVerified,
			},
		}); err != nil {
			ctx.Status(500).JSON(fiber.Map{
				"success": false, "message": "Error Updating user data",
			})
			return true, 500
		}
		dbDuration := time.Since(updateStart)
		prometheus.RecordDatabaseQueryOperation(dbDuration, "users", "UpdateOne")
	}

	err := cookies.CreateSession(models.UserRes{Username: userData.Username, ID: userData.ID, EmailVerified: true}, ctx)
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"success": true,
			"message": "Internal error",
		})

		return true, 500
	}

	return false, 200
}

func createUser(p configs.GooglePayload, collection *mongo.Collection, ctx *fiber.Ctx) (err error) {

	username := generateUsername(p.Email)
	data, err := collection.InsertOne(ctx.Context(), models.User{
		Email:         p.Email,
		Username:      username,
		Avatar:        p.Picture,
		EmailVerified: p.EmailVerified,
		Password:      "",
		CreatedAt:     time.Now(),
	})

	if err != nil {
		log.Println("Insert error:", err)
		return err
	}

	prometheus.RecordUserRegistration()
	err = cookies.CreateSession(models.UserRes{Username: username, ID: data.InsertedID.(primitive.ObjectID), EmailVerified: p.EmailVerified}, ctx)
	if err != nil {
		return err
	}

	return nil

}

func generateUsername(email string) string {
	name := strings.Split(email, "@")[0]

	var randNumbers string

	for range 5 {
		randNumbers += strconv.Itoa(rand.Intn(10))
	}

	return name + randNumbers
}
