package auth_controllers

import (
	"log"
	"main/cookies"
	"main/db"
	"main/models"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payload struct {
	Email string `json:"email"`
}

func HandleOauth(ctx *fiber.Ctx) error {
	var body Payload

	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error parsing body,", err)
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	collection := db.GetClient().Collection("users")

	var userData models.User
	if err := collection.FindOne(ctx.Context(), bson.M{"email": body.Email}).Decode(&userData); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			err := createUser(body, collection, ctx)
			if err != nil {
				log.Println("error creating user: ", err)
				return ctx.Status(500).JSON(fiber.Map{
					"success": false, "message": "Internal error",
				})
			}

			return ctx.Status(201).JSON(fiber.Map{
				"success": true,
				"message": "User created successfully",
			})
		}
		log.Printf("Database error when finding user with email %s: %v\n", body.Email, err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error getting user data",
		})
	}

	err := cookies.CreateSession(models.UserRes{Username: userData.Username, ID: userData.ID}, ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": true,
			"message": "Internal error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Welcome " + userData.Username,
	})
}

func createUser(p Payload, collection *mongo.Collection, ctx *fiber.Ctx) (err error) {

	username := generateUsername(p.Email)
	data, err := collection.InsertOne(ctx.Context(), models.User{
		Email:     p.Email,
		Username:  username,
		Password:  "",
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Insert error:", err)
		return err
	}
	err = cookies.CreateSession(models.UserRes{Username: username, ID: data.InsertedID}, ctx)
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

	log.Println(name + randNumbers)
	return name + randNumbers
}
