package task_controllers

import (
	"fmt"
	"log"
	redis "main/configs"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTaskByIdController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Id",
		})
	}

	user := ctx.Locals("user").(*models.UserRes)
	var task models.Task

	cacheKey := fmt.Sprintf("user:%s:task:%s", user.ID.Hex(), id)
	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &task, cacheKey, "task") {
		return nil
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	err = collection.FindOne(ctx.Context(), bson.M{"_id": objectId, "userId": user.ID}).Decode(&task)
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

	if err := redisClient.SetCacheData(cacheKey, ctx.Context(), task); err != nil {
		log.Println("Error setting cache: ", err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "task": task,
	})

}
