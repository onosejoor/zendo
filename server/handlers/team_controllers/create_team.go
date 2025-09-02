package team_controllers

import (
	"log"
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

	teamMember := models.TeamMemberSchema{
		TeamID: *teamID,
		UserID: user.ID,
		Role:   "owner",
	}
	if _, err := teamMember.CreateTeamMember(ctx.Context()); err != nil {
		log.Println("Error creating team_member: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to create team_member, error: " + err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"success": true, "team_id": teamID.Hex(),
	})
}
