package auth_controllers

import (
	"log"
	"main/cookies"
	"main/db"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SigninPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func HandleSignin(ctx *fiber.Ctx) error {
	var body SigninPayload

	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error parsing body,", err)
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	if err := utils.Validate.Struct(body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
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

	if !userData.ComparePassword(body.Password) {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Incorrect credentials",
		})
	}

	err := cookies.CreateSession(models.UserRes{Username: userData.Username, ID: userData.ID, EmailVerified: userData.EmailVerified}, ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success":        true,
		"message":        "Welcome " + userData.Username,
		"email_verified": userData.EmailVerified,
	})
}
