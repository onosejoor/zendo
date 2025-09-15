package team_controllers

import (
	"context"
	"fmt"
	"log"
	"main/configs/redis"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CancelInviteController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := ctx.Locals("teamId").(primitive.ObjectID)
	email := ctx.Params("email")

	if teamExist := models.CheckTeamExist(teamId, ctx.Context()); !teamExist {
		return ctx.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Team does not exist",
		})
	}

	err := models.DeleteInvite(email, teamId, ctx.Context())
	if err != nil {
		log.Println("ERROR REMOVING TEAM INVITE: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	redis.DeleteKeysByPattern(context.Background(), fmt.Sprintf("users:%s:teams:%s:invitees", user.ID.String(), teamId.String()), user.ID.String())

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Member Removed Successfully",
	})

}
