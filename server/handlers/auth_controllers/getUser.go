package auth_controllers

import (
	"fmt"
	"log"
	"main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleGetUser(ctx *fiber.Ctx) error {
	client := db.GetClient()

	var user models.User

	userCtx := ctx.Locals("user").(*models.UserRes)

	cacheKey := fmt.Sprintf("user:%s", user.ID.Hex())
	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &user, cacheKey, "user") {
		return nil
	}

	err := client.Collection("users").FindOne(ctx.Context(), bson.M{"_id": userCtx.ID}).Decode(&user)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "User does not exist",
			})
		}
		log.Printf("Database error when finding user with id %s: %v\n", userCtx.ID, err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error getting user data",
		})

	}

	_ = redisClient.SetCacheData(cacheKey, ctx.Context(), user)
	ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"user":    user,
	})
	return nil

}
