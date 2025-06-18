package project_controllers

import (
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProjectByIdController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Id",
		})
	}

	user := ctx.Locals("user").(*models.UserRes)

	client := db.GetClient()
	collection := client.Collection("tasks")

	var project models.Project
	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectId, "userId": user.ID}).Decode(&project)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Project not found",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error, try again",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "project": project,
	})

}
