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

	inviteDoc := models.CheckIfInviteExists(ctx.Context(), teamMember.Email, teamMember.TeamID)
	if inviteDoc == nil {
		return ctx.Status(404).JSON(fiber.Map{
			"success": false, "message": "Your Invite Has Expired, or It has been cancelled by the team owner",
		})
	}

	userDoc, err := models.GetUser(user.ID, ctx.Context())
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to verify user identity",
		})
	}

	// Compare emails
	if userDoc.Email != teamMember.Email {
		return ctx.Status(403).JSON(fiber.Map{
			"success": false,
			"message": "This invite was not sent to your account. Please sign in with the invited email.",
		})
	}

	teamMember.UserID = user.ID

	_, err = teamMember.CreateTeamMember(ctx.Context())
	if err != nil {
		log.Println("Error creating team_member: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to create team_member: " + err.Error(),
		})
	}

	go func(email string, teamID primitive.ObjectID) {
		if err := models.DeleteMemberInvite(context.Background(), email, teamID); err != nil {
			log.Println("Error removing member invite:", err)
		}
	}(teamMember.Email, teamMember.TeamID)

	redis.ClearAllCache(ctx.Context(), teamMember.UserID.Hex())
	redis.ClearTeamMembersCache(ctx.Context(), teamMember.TeamID)

	return ctx.Status(201).JSON(fiber.Map{
		"success": true, "team_id": teamMember.TeamID.Hex(), "message": "Team member created successfully",
	})
}
