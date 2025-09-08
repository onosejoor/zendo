package team_controllers

import (
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func DeleteTeamController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := utils.HexToObjectID(ctx.Params("id"))

	if err := models.DeleteTeam(ctx.Context(), teamId, user.ID); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete team: " + err.Error(),
		})
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex())
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "Team deleted successfully",
	})
}
