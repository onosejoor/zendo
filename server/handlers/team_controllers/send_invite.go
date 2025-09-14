package team_controllers

import (
	"context"
	"errors"
	"log"
	"main/configs/cron"
	"main/cookies"
	"main/models"
	"main/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendTeamInvite(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := ctx.Locals("teamId").(primitive.ObjectID)
	var invitePayload TeamInvite

	if err := utils.ParseBodyAndValidateStruct(&invitePayload, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	invitePayload.Email = strings.ToLower(invitePayload.Email)

	team, err := models.GetTeamById(ctx.Context(), teamId, user.ID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(404).JSON(fiber.Map{"success": false, "message": "No Team with this ID exists"})
		}
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Internal Server Error, try again"})
	}

	_, err = models.GetTeamMemberByEmail(ctx.Context(), teamId, invitePayload.Email)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Error checking membership: " + err.Error()})
	}

	invite := models.CheckIfInviteExists(ctx.Context(), invitePayload.Email, teamId)
	if invite != nil {
		if invite.Status == "sent" {
			return ctx.Status(400).JSON(fiber.Map{"success": false, "message": "Invite already exists for this email"})
		}
	}

	teamMember := &models.TeamMemberSchema{
		TeamID: teamId,
		Role:   invitePayload.Role,
		Email:  invitePayload.Email,
	}

	token := ""
	if invite == nil {
		token, err = cookies.GenerateTeamInviteEmailToken(teamMember)
		if err != nil {
			log.Println("Error generating invite token:", err)
			return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Failed to generate invite token"})
		}
	} else {
		token = invite.Token
	}

	inviteSchemaProps := &models.TeamInviteSchema{
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
			TeamName: team.Name,
		}
		emailBody := cookies.GenerateMagicTeamInviteLink(inviteProps)

		status := "sent"
		if err := cron.SendEmailToGmail(invitePayload.Email, "Invitation To Collaborate With A Team On Zendo", emailBody); err != nil {
			log.Println("Error sending invite email:", err)
			status = "failed"
		}

		inviteSchemaProps.Status = status
		_ = inviteSchemaProps.CreateOrUpdateInvite(context.Background())
	}()

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Invite created successfully. Email will be sent shortly.",
	})
}
