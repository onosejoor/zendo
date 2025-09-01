package team_controllers

import (
	"log"
	"main/models"
	"main/utils"
	"os/user"

	"github.com/gofiber/fiber/v2"
)

type TeamInvite struct {
	Email     string `json:"email" validate:"required,email"`
	TeamID    string `json:"team_id" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=member admin"`
	TeammName string `json:"team_name" validate:"required"`
}

func SendTeamInvite(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var invitePayload TeamInvite

	if err := utils.ParseBodyAndValidateStruct(&invitePayload, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

}

func CreateTeamMemberController(ctx *fiber.Ctx) error {
	var teamMember models.TeamMemberSchema

	if err := utils.ParseBodyAndValidateStruct(&teamMember, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	teamMemberId, err := teamMember.CreateTeamMember(ctx.Context())
	if err != nil {
		log.Println("Error creating team_member: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to create team_member, error: " + err.Error(),
		})
	}

	return ctx.Status(201).JSON(fiber.Map{
		"success": true, "team_id": teamMemberId.Hex(),
	})
}
