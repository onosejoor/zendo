package team_controllers

import (
	"log"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func RemoveTeamMemberController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	memberUserId := utils.HexToObjectID(ctx.Params("memberId"))
	teamId := utils.HexToObjectID(ctx.Params("teamId"))

	if user.ID == memberUserId {
		return ctx.Status(403).JSON(fiber.Map{
			"success": false, "message": "You cannot remove yourself",
		})
	}

	isOwnerOrAdmin := models.CheckMemberRoleMatch(user.ID, teamId, ctx.Context(), []string{"owner", "admin"})
	if !isOwnerOrAdmin {
		return ctx.Status(403).JSON(fiber.Map{
			"success": false, "message": "Only Admins or Owner can remove members",
		})
	}

	isMember := models.CheckMemberRoleMatch(memberUserId, teamId, ctx.Context(), []string{"member", "admin"})
	if !isMember {
		return ctx.Status(403).JSON(fiber.Map{
			"success": false, "message": "Only Admins or Member can be removed or member does not exist",
		})
	}

	err := models.DeleteTeamMember(memberUserId, teamId, ctx.Context())
	if err != nil {
		log.Println("ERROR REMOVING TEAM MEMBER FROM TEAM: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	redis.DeleteTeamsCache(ctx.Context(), user.ID.Hex())
	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Member Removed Successfully",
	})

}
