package task_controllers

import (
	"errors"
	"fmt"
	"log"
	prometheus "main/configs/prometheus"
	redis "main/configs/redis"
	"main/db"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTaskByIdController(ctx *fiber.Ctx) error {

	taskId := utils.HexToObjectID(ctx.Params("id"))

	user := ctx.Locals("user").(*models.UserRes)
	var task models.Task

	cacheKey := fmt.Sprintf("user:%s:tasks:%s", user.ID.Hex(), taskId.Hex())
	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &task, cacheKey, "task") {
		prometheus.RecordRedisOperation("get_cache")
		return nil
	}
	prometheus.RecordRedisOperation("cache_miss")
	client := db.GetClient()
	collection := client.Collection("tasks")

	err := collection.FindOne(ctx.Context(), bson.M{"_id": taskId, "userId": user.ID, "team_id": nil}).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
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
	prometheus.RecordRedisOperation("set_cache")
	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "task": task,
	})

}
