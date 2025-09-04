package team_task_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTeamTasksController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	teamID := utils.HexToObjectID(ctx.Params("teamId"))

	var teamsSlice = make([]models.Task, 0)

	cacheKey := fmt.Sprintf("user:%s:teams:%s:tasks:page:%d:limit:%d", user.ID.Hex(), teamID.Hex(), page, limit)

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &teamsSlice, cacheKey, "tasks") {
		return nil
	}

	teamSlice, err := models.GetTasksForTeam(ctx.Context(), teamID, page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch tasks: " + err.Error(),
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), teamSlice)

	return ctx.JSON(fiber.Map{
		"success": true,
		"tasks":   teamSlice,
	})

}
