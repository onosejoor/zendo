package team_controllers

import (
	"context"
	"errors"
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

	err := cancelInvite(teamId, inviteId)
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

func cancelInvite(teamId, inviteId primitive.ObjectID) error {
	if !models.CheckTeamExist(teamId, context.TODO()) {
		return errors.New("team not found")
	}
	return models.DeleteInvite(inviteId, teamId, context.TODO())
}
