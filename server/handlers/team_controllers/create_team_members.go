package team_controllers

import (
	"log"
	"main/configs/redis"
	"main/cookies"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

type TeamInvite struct {
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,oneof=member admin"`
	TeamName string `json:"team_name" validate:"required"`
}

func CreateTeamMemberController(ctx *fiber.Ctx) error {
	inviteToken := ctx.Query("token")
	if inviteToken == "" {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Token is required",
		})
	}

	claims, err := cookies.VerifyTeamEmailInviteToken(inviteToken)
	if err != nil {
		log.Println("Error verifying invite token: ", err)
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Invalid Invite Token, or Token is Expired",
		})
	}
	teamMember := claims.TeamMemberSchema

	_, err = teamMember.CreateTeamMember(ctx.Context())
	if err != nil {
		log.Println("Error creating team_member: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to create team_member: " + err.Error(),
		})
	}

	go func() {
		err = models.DeleteMemberInvite(ctx.Context(), teamMember.Email, teamMember.TeamID)
		if err != nil {
			log.Println("Error removing member invite: ", err)
		}
	}()

	redis.ClearAllCache(ctx.Context(), teamMember.UserID.Hex())
	return ctx.Status(201).JSON(fiber.Map{
		"success": true, "team_id": teamMember.TeamID.Hex(), "message": "Team member created successfully",
	})
}
