package project_controllers

import (
	"fmt"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetProjectByIdController(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	objectId, _ := primitive.ObjectIDFromHex(id)

	user := ctx.Locals("user").(*models.UserRes)
	var project models.Project

	redisClient := redis.GetRedisClient()
	cacheKey := fmt.Sprintf("user:%s:project:%s", user.ID.Hex(), id)

	if redisClient.GetCacheHandler(ctx, &project, cacheKey, "project") {
		return nil
	}

	client := db.GetClient()
	collection := client.Collection("projects")

	err := collection.FindOne(ctx.Context(), bson.M{"_id": objectId, "ownerId": user.ID}).Decode(&project)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Project not found",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error, try again",
		})
	}

	if err := redisClient.SetCacheData(cacheKey, ctx.Context(), project); err != nil {
		return nil
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "project": project,
	})

}
