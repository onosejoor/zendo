package project_controllers

import (
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteProjectController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := ctx.Locals("user").(*models.UserRes)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Id",
		})
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	if err := collection.FindOneAndDelete(ctx.Context(), bson.M{"_id": objectId, "userId": user.ID}).Err(); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Project not found",
			})
		}
		log.Println("Error deleting project: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error",
		})
	}

	if err := redis.DeleteTaskCache(ctx.Context(), user.ID.Hex(), id); err != nil {
		log.Println(err.Error())
	}
	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Project deleted",
	})
}
