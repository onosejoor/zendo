package team_controllers

import (
	"log"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RemoveTeamMemberController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	memberUserId := utils.HexToObjectID(ctx.Params("memberId"))
	teamId := ctx.Locals("teamId").(primitive.ObjectID)

	if teamExist := models.CheckTeamExist(teamId, ctx.Context()); !teamExist {
		return ctx.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Team does not exist",
		})
	}

	isMember := models.CheckMemberRoleMatch(memberUserId, teamId, ctx.Context(), []string{"member", "admin"})
	if !isMember {
		return ctx.Status(403).JSON(fiber.Map{
			"success": false, "message": "Only Admins or Member can be removed  or member does not exist",
		})
	}

	err := models.DeleteTeamMember(memberUserId, teamId, ctx.Context())
	if err != nil {
		log.Println("ERROR REMOVING TEAM MEMBER FROM TEAM: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	go redis.DeleteTeamsCache(ctx.Context(), user.ID.Hex())
	go redis.DeleteTeamsCache(ctx.Context(), memberUserId.Hex())
	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Member Removed Successfully",
	})

}
