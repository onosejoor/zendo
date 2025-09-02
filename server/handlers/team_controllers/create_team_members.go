package team_controllers

import (
	"context"
	"log"
	"main/configs/cron"
	"main/cookies"
	"main/db"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamInvite struct {
	Email     string `json:"email" validate:"required,email"`
	TeamID    string `json:"team_id" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=member admin"`
	TeammName string `json:"team_name" validate:"required"`
}

func teamMemberExist(ctx context.Context, invitePayload *TeamInvite) (bool, primitive.ObjectID) {
	client := db.GetClient().Collection("team_members")

	var teamMember models.TeamMemberSchema
	err := client.FindOne(ctx, bson.M{
		"email":   invitePayload.Email,
		"team_id": invitePayload.TeamID,
	}).Decode(&teamMember)
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		return true, primitive.NilObjectID
	}
	return false, teamMember.UserID
}

func SendTeamInvite(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var invitePayload TeamInvite

	if err := utils.ParseBodyAndValidateStruct(&invitePayload, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	teamMemberExist, userId := teamMemberExist(ctx.Context(), &invitePayload)
	if teamMemberExist {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "User is already a member of the team",
		})
	}

	teamMember := &models.TeamMemberSchema{
		TeamID: utils.HexToObjectID(invitePayload.TeamID),
		UserID: userId,
		Role:   invitePayload.Role,
	}

	token, err := cookies.GenerateTeamInviteEmailToken(teamMember)
	if err != nil {
		log.Println("Error generating invite token: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to generate invite token",
		})
	}
	inviteProps := cookies.InviteTokenProps{
		Token:    token,
		Role:     invitePayload.Role,
		Username: user.Username,
		TeamName: invitePayload.TeammName,
	}

	emailBody := cookies.GenerateMagicTeamInviteLink(inviteProps)

	if err := cron.SendEmailToGmail(invitePayload.Email, "Invitation To Collabprate With A Team On Zendo", emailBody); err != nil {
		log.Println("Error sending invite email: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to send invite email",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Invite sent successfully",
	})

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
			"success": false, "message": "Invalid Invite Token, or Toke is Expired",
		})
	}
	teamMember := claims.TeamMemberSchema

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
