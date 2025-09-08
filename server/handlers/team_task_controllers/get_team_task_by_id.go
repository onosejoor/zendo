package team_task_controllers

import (
	"errors"
	"fmt"
	"main/configs/redis"
	"main/db"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetTeamTaskByIdController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)

	teamID := utils.HexToObjectID(ctx.Params("teamId"))
	taskId := utils.HexToObjectID(ctx.Params("id"))

	var teamsTask models.Task

	cacheKey := fmt.Sprintf("user:%s:teams:%s:tasks:%s", user.ID.Hex(), teamID.Hex(), taskId.Hex())

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &teamsTask, cacheKey, "task") {
		return nil
	}

	taskCollection := db.GetClient().Collection("tasks")

	err := taskCollection.FindOne(ctx.Context(), bson.M{
		"_id":     taskId,
		"team_id": teamID,
	}).Decode(&teamsTask)

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

	redisClient.SetCacheData(cacheKey, ctx.Context(), teamsTask)

	return ctx.JSON(fiber.Map{
		"success": true,
		"task":    teamsTask,
	})

}
