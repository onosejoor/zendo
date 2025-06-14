package handlers

import (
	"log"
	"main/auth/cookies"
	"main/auth/models"
	"main/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleSignup(ctx *fiber.Ctx) error {
	var body models.UserPayload

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	collection := db.GetClient().Collection("users")
	count, err := collection.CountDocuments(ctx.Context(), bson.M{
		"email": body.Email,
	})
	if err != nil {
		log.Println("DB error:", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	if count > 0 {
		return ctx.Status(409).JSON(fiber.Map{
			"success": false,
			"message": "User already exists",
		})
	}

	id, err := models.CreateUser(body, collection, ctx.Context())
	if err != nil {
		log.Println("error creating user: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error",
		})
	}

	err = cookies.CreateSession(models.UserRes{Username: body.Username, ID: id}, ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": true,
			"message": "Internal error",
		})
	}

	ctx.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
	})
	return nil
}
