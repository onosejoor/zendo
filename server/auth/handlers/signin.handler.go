package handlers

import (
	"log"
	"main/auth/cookies"
	"main/auth/models"
	"main/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func HandleSignin(ctx *fiber.Ctx) error {
	var body models.UserPayload

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
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "User not found",
			})
		}
		log.Printf("Database error when finding user with email %s: %v\n", body.Email, err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error getting user data",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(body.Password)); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Incorrect Password",
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
