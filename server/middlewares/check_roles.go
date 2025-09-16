package middlewares

import (
	"main/models"
	"main/utils"
	"slices"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RequireTeamMember(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*models.UserRes)
		teamId := utils.HexToObjectID(c.Params("teamId"))

		if teamId == primitive.NilObjectID {
			return c.Status(400).JSON(fiber.Map{
				"success": false,
				"message": "Team not found",
			})
		}

		exists, role, err := models.IsTeamMembers(c.Context(), []primitive.ObjectID{user.ID}, teamId, false)
		if err != nil {
			return c.Status(403).JSON(fiber.Map{
				"success": false,
				"message": "error performing tasks",
			})
		}

		if !exists {
			return c.Status(403).JSON(fiber.Map{
				"success": false,
				"message": "Team not found or You are not a member of this team",
			})
		}

		if len(roles) > 0 && !slices.Contains(roles, role) {
			return c.Status(403).JSON(fiber.Map{
				"success": false,
				"message": "You are not authorized to perform this action",
			})
		}

		c.Locals("teamId", teamId)
		c.Locals("role", role)

		return c.Next()
	}
}
