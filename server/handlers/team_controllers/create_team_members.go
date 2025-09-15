package team_controllers

import (
	"context"
	"log"
	"main/configs/redis"
	"main/cookies"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamInvite struct {
	Email string `json:"email" validate:"required,email"`
	Role  string `json:"role" validate:"required,oneof=member admin"`
}

func CreateTeamMemberController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)

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

	if teamMember.UserID == primitive.NilObjectID {
		teamMember.UserID = user.ID
	}

	inviteDoc := models.CheckIfInviteExists(ctx.Context(), teamMember.Email, teamMember.TeamID)
	if inviteDoc == nil {
		return ctx.Status(404).JSON(fiber.Map{
			"success": false, "message": "Your Invite Has Expired, or It has been councelled by the team owner",
		})
	}

	_, err = teamMember.CreateTeamMember(ctx.Context())
	if err != nil {
		log.Println("Error creating team_member: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to create team_member: " + err.Error(),
		})
	}

	go func() {
		err = models.DeleteMemberInvite(context.Background(), teamMember.Email, teamMember.TeamID)
		if err != nil {
			log.Println("Error removing member invite: ", err)
		}
	}()

	redis.ClearAllCache(ctx.Context(), teamMember.UserID.Hex())
	go redis.ClearTeamMembersCache(context.Background(), teamMember.TeamID)
	return ctx.Status(201).JSON(fiber.Map{
		"success": true, "team_id": teamMember.TeamID.Hex(), "message": "Team member created successfully",
	})
}
