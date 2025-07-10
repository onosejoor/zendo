package project_controllers

import (
	"fmt"
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

	objectId, _ := primitive.ObjectIDFromHex(id)

	client := db.GetClient()
	projectCollection := client.Collection("projects")
	taskCollection := client.Collection("tasks")

	if err := projectCollection.FindOneAndDelete(ctx.Context(), bson.M{"_id": objectId, "ownerId": user.ID}).Err(); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false,
				"message": "Project not found",
			})
		}
		log.Println("Error deleting project: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal error",
		})
	}

	_, err := taskCollection.DeleteMany(ctx.Context(), bson.M{
		"projectId": objectId,
		"userId":    user.ID,
	})
	if err != nil {
		log.Println("Error deleting related tasks: ", err.Error())
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex(), "", id)

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Project and related tasks deleted",
	})
}

func DeleteAllProjectsController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)

	deletedData, err := models.DeleteAllProjectsWithTransaction(ctx.Context(), user)
	if err != nil {
		log.Println("Error deleing all tasks: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error deleting all tasks, try again",
		})
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex(), "", "")

	message := fmt.Sprintf("Deleted %v tasks and %v project(s)", deletedData.TotalTasksDeleted, deletedData.TotalProjectsDeleted)

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": message,
	})

}
