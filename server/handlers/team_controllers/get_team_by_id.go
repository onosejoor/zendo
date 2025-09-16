package team_controllers

import (
	"errors"
	"fmt"
	"main/configs/redis"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTeamByIDController(ctx *fiber.Ctx) error {
	teamId := ctx.Locals("teamId").(primitive.ObjectID)
	user := ctx.Locals("user").(*models.UserRes)

	var team *models.TeamWithMemberAndRole

	cacheKey := fmt.Sprintf("user:%s:teams:%s", user.ID.Hex(), teamId.Hex())

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &team, cacheKey, "team") {
		return nil
	}

	team, err := models.GetTeamById(ctx.Context(), teamId, user.ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Team Does not exist",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch team: " + err.Error(),
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), *team)

	return ctx.JSON(fiber.Map{
		"success": true,
		"team":    team,
	})

}
