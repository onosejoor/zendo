package team_controllers

import (
	"main/configs/redis"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTeamController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := ctx.Locals("teamId").(primitive.ObjectID)

	if err := models.DeleteTeam(ctx.Context(), teamId, user.ID); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex())
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Team deleted successfully",
	})
}
