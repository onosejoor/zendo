package project_controllers

import (
	"fmt"
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProjectsController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var projects = make([]models.Project, 0)

	cacheKey := fmt.Sprintf("user:%s:projects", user.ID.Hex())
	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &projects, cacheKey, "projects") {
		return nil
	}

	client := db.GetClient()
	collection := client.Collection("projects")

	cursor, err := collection.Find(ctx.Context(), bson.M{"ownerId": user.ID})
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	if err := cursor.All(ctx.Context(), &projects); err != nil {
		log.Println("Error parsing db data to slice: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	err = redisClient.SetCacheData(cacheKey, ctx.Context(), projects)
	if err != nil {
		log.Println("Error caching data: ", err.Error())
	}

	return ctx.Status(200).JSON(bson.M{
		"success":  true,
		"projects": projects,
	})

}
