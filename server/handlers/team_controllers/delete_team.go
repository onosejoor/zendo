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

	if teamExist := models.CheckTeamExist(teamId, ctx.Context()); !teamExist {
		return ctx.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Team does not exist",
		})
	}

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
