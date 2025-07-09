package task_controllers

import (
	"fmt"
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllTasksController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var dbTaks = make([]models.Task, 0)

	cacheKey := fmt.Sprintf("user:%s:tasks", user.ID.Hex())
	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &dbTaks, cacheKey, "tasks") {
		return nil
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	cursor, err := collection.Find(ctx.Context(), bson.M{"userId": user.ID}, &options.FindOptions{
		Sort: bson.M{"created_at": -1},
	})
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	if err := cursor.All(ctx.Context(), &dbTaks); err != nil {
		log.Println("Error parsing db data to slice: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	err = redisClient.SetCacheData(cacheKey, ctx.Context(), dbTaks)
	if err != nil {
		log.Println("Error caching data: ", err.Error())
	}

	return ctx.Status(200).JSON(bson.M{
		"success": true,
		"tasks":   dbTaks,
	})

}
