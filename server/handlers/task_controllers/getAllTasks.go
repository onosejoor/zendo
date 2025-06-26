package task_controllers

import (
	"encoding/json"
	"fmt"
	"log"
	redis "main/configs"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasksController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var dbTaks = make([]models.Task, 0)

	cacheKey := fmt.Sprintf("user:%s:task", user.ID.Hex())
	redisClient := redis.GetRedisClient()

	data, err, isEmpty := redisClient.GetCacheData(cacheKey, ctx.Context())
	if err != nil && !isEmpty {
		log.Println("Error querying casched data: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	if !isEmpty {
		_ = json.Unmarshal(data, &dbTaks)
		return ctx.Status(200).JSON(bson.M{
			"success": true,
			"tasks":   dbTaks,
		})
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	cursor, err := collection.Find(ctx.Context(), bson.M{"userId": user.ID})
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

	jsonValue, _ := json.Marshal(&dbTaks)
	err = redisClient.SetCacheData(cacheKey, ctx.Context(), string(jsonValue))
	if err != nil {
		log.Println("Error storing data in cache: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	return ctx.Status(200).JSON(bson.M{
		"success": true,
		"tasks":   dbTaks,
	})

}
