package team_controllers

import (
	"context"
	"errors"
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

const (
	USER_DOES_NOT_EXIST = "user does not exist"
)

func teamMemberExist(ctx context.Context, invitePayload *TeamInvite, teamId primitive.ObjectID) (primitive.ObjectID, bool, error) {
	usersColl := db.GetClient().Collection("users")
	membersColl := db.GetClient().Collection("team_members")

	var user models.GetUserByEmailPayload
	user, err := models.GetUserByEmail(invitePayload.Email, usersColl, ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, false, errors.New(USER_DOES_NOT_EXIST)
		}
		return primitive.NilObjectID, false, err
	}

	var teamMember models.TeamMemberSchema
	err = membersColl.FindOne(ctx, bson.M{
		"user_id": user.ID,
		"team_id": teamId,
	}).Decode(&teamMember)

	if err == nil {
		return user.ID, true, nil
	}
	if err != mongo.ErrNoDocuments {
		return primitive.NilObjectID, false, err
	}

	return user.ID, false, nil
}

func SendTeamInvite(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := utils.HexToObjectID(ctx.Params("id"))
	var invitePayload TeamInvite

	if err := utils.ParseBodyAndValidateStruct(&invitePayload, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	userId, exists, err := teamMemberExist(ctx.Context(), &invitePayload, teamId)
	if err != nil {
		if err.Error() == USER_DOES_NOT_EXIST {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "User with this email does not exist",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error checking team membership: " + err.Error(),
		})
	}

	if exists {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "User is already a member of the team",
		})
	}

	teamMember := &models.TeamMemberSchema{
		TeamID: teamId,
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
		TeamName: invitePayload.TeamName,
	}

	emailBody := cookies.GenerateMagicTeamInviteLink(inviteProps)

	if err := cron.SendEmailToGmail(invitePayload.Email, "Invitation To Collaborate With A Team On Zendo", emailBody); err != nil {
		log.Println("Error sending invite email: ", err)

		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Failed to send invite email",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Invite sent successfully",
	})

}
