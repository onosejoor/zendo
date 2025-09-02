package routes

import (
	"main/handlers/team_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TeamsRoutes(app fiber.Router) {
	teamsRoute := app.Group("/teams")
	teamsRoute.Use(middlewares.AuthMiddleware)

	teamsRoute.Get("/", team_controllers.GetTeamsController)
	teamsRoute.Post("/", team_controllers.CreateTeamController)
	teamsRoute.Get("/team_members", team_controllers.GetTeamMembersController)
	teamsRoute.Post("/team_members/invite", team_controllers.SendTeamInvite)
	teamsRoute.Post("/team_members/invite/callback", team_controllers.CreateTeamMemberController)

}
