package project_controllers

import (
	"log"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTaskInProjectsController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Id",
		})
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	user := ctx.Locals("user").(*models.UserRes)

	cursor, err := collection.Find(ctx.Context(), bson.M{"userId": user.ID, "projectId": objectId})
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	var tasks []models.Task

	if err := cursor.All(ctx.Context(), &tasks); err != nil {
		log.Println("Error parsing db data to slice: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	return ctx.Status(200).JSON(bson.M{
		"success": true,
		"tasks":   tasks,
	})

}
