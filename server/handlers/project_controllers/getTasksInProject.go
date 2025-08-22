package project_controllers

import (
	"fmt"
	"log"
	prometheus "main/configs/prometheus"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		prometheus.RecordRedisOperation("get_cache")
		return nil
	}
	prometheus.RecordRedisOperation("clear_cache")
	client := db.GetClient()
	collection := client.Collection("tasks")

	cursor, err := collection.Find(ctx.Context(), bson.M{"userId": user.ID, "projectId": objectId}, &options.FindOptions{
		Sort: bson.M{"created_at": -1},
	})
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
	prometheus.RecordRedisOperation("set_cache")
	return ctx.Status(200).JSON(bson.M{
		"success": true,
		"tasks":   tasks,
	})

}
