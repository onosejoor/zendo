package project_controllers

import (
	"fmt"
	"log"
	prometheus "main/configs/prometheus"
	redis "main/configs/redis"
	"main/db"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProjectsController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var projects = make([]models.Project, 0)

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	cacheKey := fmt.Sprintf("user:%s:projects:page:%s:limit:%s", user.ID.Hex(), page, limit)
	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &projects, cacheKey, "projects") {
		prometheus.RecordRedisOperation("get_cache")
		return nil
	}
	prometheus.RecordRedisOperation("cache_miss")
	client := db.GetClient()
	collection := client.Collection("projects")

	opts := utils.GeneratePaginationOptions(page, limit)
	cursor, err := collection.Find(ctx.Context(), bson.M{"ownerId": user.ID}, opts)
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
	prometheus.RecordRedisOperation("set_cache")
	if err != nil {
		log.Println("Error caching data: ", err.Error())
	}

	return ctx.Status(200).JSON(bson.M{
		"success":  true,
		"projects": projects,
	})

}
