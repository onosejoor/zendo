package team_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTeamByIdStatsController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	role := ctx.Locals("role").(string)
	teamId := ctx.Locals("teamId").(primitive.ObjectID)

	cacheKey := fmt.Sprintf("user:%s:teams:%s:stats", user.ID.Hex(), teamId.Hex())

	var stats *models.TeamStatsRes

	redisClient := redis.GetRedisClient()
	if redisClient.GetCacheHandler(ctx, &stats, cacheKey, "stats") {
		return nil
	}

	stats, err := models.GetTeamByIdStats(teamId, ctx.Context(), role)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error Getting Stats, Try Again",
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), *stats)

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "stats": *stats,
	})

}
