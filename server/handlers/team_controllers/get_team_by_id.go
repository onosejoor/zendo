package team_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTeamByIDController(ctx *fiber.Ctx) error {
	teamId := utils.HexToObjectID(ctx.Params("id"))
	user := ctx.Locals("user").(*models.UserRes)

	var team *models.TeamSchema

	cacheKey := fmt.Sprintf("user:%s:teams:%s", user.ID.Hex(), teamId.Hex())

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &team, cacheKey, "team") {
		return nil
	}

	team, err := models.GetTeamById(ctx.Context(), teamId, user.ID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
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
