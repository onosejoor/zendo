package team_task_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

type GetTeamTasksResponse struct {
	Tasks []models.TaskWithAssignees `json:"tasks"`
	Role  string                     `json:"role"`
}

func GetTeamTasksController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	role := ctx.Locals("role").(string)

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	teamID := utils.HexToObjectID(ctx.Params("teamId"))

	var teamsTasksRes GetTeamTasksResponse

	cacheKey := fmt.Sprintf("user:%s:teams:%s:tasks:page:%d:limit:%d", user.ID.Hex(), teamID.Hex(), page, limit)

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &teamsTasksRes, cacheKey, "data") {
		return nil
	}

	teamTasksSlice, err := models.GetTasksForTeam(ctx.Context(), teamID, page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch tasks: " + err.Error(),
		})
	}

	if len(teamTasksSlice) < 1 {
		teamTasksSlice = []models.TaskWithAssignees{}
	}

	teamsTasksRes = GetTeamTasksResponse{
		Tasks: teamTasksSlice, Role: role,
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), teamsTasksRes)

	return ctx.JSON(fiber.Map{
		"success": true,
		"data":    teamsTasksRes,
	})

}
