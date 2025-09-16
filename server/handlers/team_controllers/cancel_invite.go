package team_controllers

import (
	"context"
	"log"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CancelInviteController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := ctx.Locals("teamId").(primitive.ObjectID)
	inviteId := utils.HexToObjectID(ctx.Params("inviteId"))

	if teamExist := models.CheckTeamExist(teamId, ctx.Context()); !teamExist {
		return ctx.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Team does not exist",
		})
	}

	err := models.DeleteInvite(inviteId, teamId, ctx.Context())
	if err != nil {
		log.Println("ERROR REMOVING TEAM INVITE: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	if err := redis.DeleteTeamsCache(context.Background(), user.ID.Hex()); err != nil {
		log.Println(err)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Invite Cancelled Successfully",
	})

}
