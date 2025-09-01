package team_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func GetTeamsController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	var teamsSlice = make([]models.TeamWithRole, 0)

	cacheKey := fmt.Sprintf("user:%s:teams:page:%d:limit:%d", user.ID.Hex(), page, limit)

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &teamsSlice, cacheKey, "teams") {
		return nil
	}

	teamsSlice, err := models.GetTeamsForUser(ctx.Context(), user.ID, page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch teams, error: " + err.Error(),
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), teamsSlice)

	return ctx.JSON(fiber.Map{
		"success": true,
		"teams":   teamsSlice,
	})
}
