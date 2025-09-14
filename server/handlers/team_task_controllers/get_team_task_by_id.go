package team_task_controllers

import (
	"errors"
	"fmt"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetTeamTaskByIdRes struct {
	Task models.TaskWithAssignees `json:"task"`
	Role string                   `json:"role"`
}

func GetTeamTaskByIdController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	role := ctx.Locals("role").(string)

	teamID := ctx.Locals("teamId").(primitive.ObjectID)
	taskId := utils.HexToObjectID(ctx.Params("id"))

	var response GetTeamTaskByIdRes

	cacheKey := fmt.Sprintf("user:%s:teams:%s:tasks:%s", user.ID.Hex(), teamID.Hex(), taskId.Hex())

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &response, cacheKey, "data") {
		return nil
	}

	var teamTask *models.TaskWithAssignees
	teamTask, err := models.GetTaskForTeamById(ctx.Context(), teamID, taskId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Task Not found for this team",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch task: " + err.Error(),
		})
	}

	response = GetTeamTaskByIdRes{
		Task: *teamTask, Role: role,
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), response)

	return ctx.JSON(fiber.Map{
		"success": true,
		"data":    response,
	})

}
