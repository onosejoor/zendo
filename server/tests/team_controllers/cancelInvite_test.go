package team_controllers_test

import (
	"net/http/httptest"
	"testing"

	"main/handlers/team_controllers"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCancelInviteController_NoTeam(t *testing.T) {
	utils.PullEnv()
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user", &models.UserRes{ID: primitive.NewObjectID()})
		c.Locals("teamId", primitive.NewObjectID())
		return c.Next()
	})

	app.Delete("/teams/:teamId/invites/:inviteId", team_controllers.CancelInviteController)

	// use valid ObjectID for inviteId
	validInviteId := primitive.NewObjectID().Hex()
	req := httptest.NewRequest("DELETE", "/teams/000000000000000000000000/invites/"+validInviteId, nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 404, resp.StatusCode)
}
