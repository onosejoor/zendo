package team_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTeamMembersController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := utils.HexToObjectID(ctx.Params("id"))
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	var teamMembersSlice = make([]models.UserWithRole, 0)

	cacheKey := fmt.Sprintf("user:%s:teams:%s:members:page:%d:limit:%d", user.ID.Hex(), teamId.Hex(), page, limit)

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &teamMembersSlice, cacheKey, "members") {
		return nil
	}

	teamMembersSlice, err := models.GetUsersForTeam(ctx.Context(), teamId, page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch team members, error: " + err.Error(),
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), teamMembersSlice)

	return ctx.JSON(fiber.Map{
		"success": true,
		"members": teamMembersSlice,
	})
}
