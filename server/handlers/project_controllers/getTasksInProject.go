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
)

func GetTaskInProjectsController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Id",
		})
	}

	var tasks = make([]models.Task, 0)
	user := ctx.Locals("user").(*models.UserRes)

	cacheKey := fmt.Sprintf("user:%s:project:%s:tasks", user.ID.Hex(), id)
	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &tasks, cacheKey, "tasks") {
		return nil
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	cursor, err := collection.Find(ctx.Context(), bson.M{"userId": user.ID, "projectId": objectId})
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	if err := cursor.All(ctx.Context(), &tasks); err != nil {
		log.Println("Error parsing db data to slice: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	err = redisClient.SetCacheData(cacheKey, ctx.Context(), tasks)
	if err != nil {
		log.Println("Error caching data: ", err.Error())
	}

	return ctx.Status(200).JSON(bson.M{
		"success": true,
		"tasks":   tasks,
	})

}
