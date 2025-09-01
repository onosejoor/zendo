package team_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func GetTeamMembersController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	var teamMembersSlice = make([]models.UserWithRole, 0)

	cacheKey := fmt.Sprintf("user:%s:team_members:page:%d:limit:%d", user.ID.Hex(), page, limit)

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &teamMembersSlice, cacheKey, "team_members") {
		return nil
	}

	teamMembersSlice, err := models.GetUsersForTeam(ctx.Context(), user.ID, page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch team members, error: " + err.Error(),
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), teamMembersSlice)

	return ctx.JSON(fiber.Map{
		"success":      true,
		"team_members": teamMembersSlice,
	})
}
