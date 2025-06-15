package task_controllers

import (
	"fmt"
	"log"
	"main/cookies"
	"main/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteTaskController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	cookie := ctx.Cookies("zendo_session_token")

	userId, err := cookies.VerifyJwt(cookie)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"success": false, "message": fmt.Sprintf("Error: %v", err.Error()),
		})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Id",
		})
	}
	user, err := primitive.ObjectIDFromHex(userId.(string))
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Id",
		})
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	if err := collection.FindOneAndDelete(ctx.Context(), bson.M{"_id": objectId, "userId": user}).Err(); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Task not found",
			})
		}
		log.Println("Error deleting task: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Task deleted",
	})
}
