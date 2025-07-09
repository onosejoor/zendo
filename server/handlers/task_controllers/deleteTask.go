package task_controllers

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

func DeleteTaskController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user := ctx.Locals("user").(*models.UserRes)

	objectId, _ := primitive.ObjectIDFromHex(id)

	client := db.GetClient()
	collection := client.Collection("tasks")

	var task models.Task

	err := collection.FindOne(ctx.Context(), bson.M{"_id": objectId, "userId": user.ID}).Decode(&task)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Task not found",
			})

		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error, try again",
		})
	}

	if task.ProjectId != nil {
		err := task.DeleteTaskWithTransaction(ctx.Context())
		if err != nil {
			log.Println("Error deleting task: ", err)

			return ctx.Status(500).JSON(fiber.Map{
				"success": false, "message": "Error parsing data",
			})
		}
		redis.ClearAllCache(ctx.Context(), user.ID.Hex(), task.ID.Hex(), task.ProjectId.Hex())

		return ctx.Status(200).JSON(fiber.Map{
			"success": true, "message": "task deleted",
		})
	}

	if _, err := collection.DeleteOne(ctx.Context(), bson.M{"_id": objectId, "userId": user.ID}); err != nil {
		log.Println("Error deleting task: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error",
		})
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex(), task.ID.Hex(), "")

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Task deleted",
	})
}
