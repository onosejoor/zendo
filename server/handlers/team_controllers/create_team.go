package team_controllers

import (
	"log"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTeamController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var team models.TeamSchema

	if err := utils.ParseBodyAndValidateStruct(&team, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	teamID, err := team.CreateTeam(ctx.Context(), user.ID)
	if err != nil {
		log.Println("Error creating team: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to create team, error: " + err.Error(),
		})
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex())

	return ctx.Status(201).JSON(fiber.Map{
		"success": true, "teamId": teamID.Hex(), "message": "Team Created Successfully",
	})
}
