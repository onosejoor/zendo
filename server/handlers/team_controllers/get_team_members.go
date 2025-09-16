package team_controllers

import (
	"fmt"
	"main/configs/redis"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTeamMembersResponse struct {
	Members []models.UserWithRole `json:"members"`
	Role    string                `json:"role"`
}

func GetTeamMembersController(ctx *fiber.Ctx) error {

	user := ctx.Locals("user").(*models.UserRes)
	teamId := ctx.Locals("teamId").(primitive.ObjectID)
	role := ctx.Locals("role").(string)

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	var teamMembersSlice GetTeamMembersResponse

	cacheKey := fmt.Sprintf("user:%s:teams:%s:members:page:%d:limit:%d", user.ID.Hex(), teamId.Hex(), page, limit)

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &teamMembersSlice, cacheKey, "data") {
		return nil
	}

	teamMembers, err := models.GetUsersForTeam(ctx.Context(), teamId, page, limit)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to fetch team members, error: " + err.Error(),
		})
	}

	res := GetTeamMembersResponse{
		Role:    role,
		Members: teamMembers,
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), res)

	return ctx.JSON(fiber.Map{
		"success": true,
		"data":    res,
	})
}
