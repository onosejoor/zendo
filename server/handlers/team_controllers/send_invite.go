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
	membersColl := db.GetClient().Collection("team_members")

	user, err := models.GetUserByEmail(invitePayload.Email, ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return primitive.NilObjectID, false, errors.New(USER_DOES_NOT_EXIST)
		}
		return primitive.NilObjectID, false, err
	}

	filter := bson.M{"user_id": user.ID, "team_id": teamId}
	exists := membersColl.FindOne(ctx, filter).Err() == nil

	return user.ID, exists, nil
}

func SendTeamInvite(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := utils.HexToObjectID(ctx.Params("id"))
	var invitePayload TeamInvite

	if err := utils.ParseBodyAndValidateStruct(&invitePayload, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	userId, exists, err := teamMemberExist(ctx.Context(), &invitePayload, teamId)
	if err != nil {
		if err.Error() == USER_DOES_NOT_EXIST {
			return ctx.Status(404).JSON(fiber.Map{"success": false, "message": "User with this email does not exist"})
		}
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Error checking membership: " + err.Error()})
	}
	if exists {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": "User is already a member of the team"})
	}
	if models.CheckIfInviteExists(ctx.Context(), invitePayload.Email, teamId) {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": "Invite already exists for this team"})
	}

	teamMember := &models.TeamMemberSchema{
		TeamID: teamId,
		UserID: userId,
		Role:   invitePayload.Role,
		Email:  invitePayload.Email,
	}

	token, err := cookies.GenerateTeamInviteEmailToken(teamMember)
	if err != nil {
		log.Println("Error generating invite token:", err)
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Failed to generate invite token"})
	}

	inviteSchemaProps := models.TeamInviteSchema{
		Email:  invitePayload.Email,
		TeamID: teamId,
		Token:  token,
		Status: "pending",
	}
	if err := inviteSchemaProps.CreateOrUpdateInvite(ctx.Context()); err != nil {
		log.Println("ERROR ADDING USER INVITE TO DB:", err)
	}

	go func() {
		inviteProps := cookies.InviteTokenProps{
			Token:    token,
			Role:     invitePayload.Role,
			Username: user.Username,
			TeamName: invitePayload.TeamName,
		}
		emailBody := cookies.GenerateMagicTeamInviteLink(inviteProps)

		status := "sent"
		if err := cron.SendEmailToGmail(invitePayload.Email, "Invitation To Collaborate With A Team On Zendo", emailBody); err != nil {
			log.Println("Error sending invite email:", err)
			status = "failed"
		}

		inviteSchemaProps.Status = status
		_ = inviteSchemaProps.CreateOrUpdateInvite(ctx.Context())
	}()

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Invite created successfully. Email will be sent shortly.",
	})
}
