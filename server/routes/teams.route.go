package routes

import (
	"main/handlers/team_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TeamsRoutes(app fiber.Router) {
	teamsRoute := app.Group("/teams")
	app.Post("/teams/members/invite/callback", team_controllers.CreateTeamMemberController)

	teamsRoute.Use(middlewares.AuthMiddleware)

	teamsRoute.Get("/", team_controllers.GetTeamsController)
	teamsRoute.Get("/:id", team_controllers.GetTeamByIDController)
	teamsRoute.Post("/", team_controllers.CreateTeamController)
	teamsRoute.Get("/:id/members", team_controllers.GetTeamMembersController)
	teamsRoute.Post("/:id/members/invite", team_controllers.SendTeamInvite)
	teamsRoute.Delete("/:id", team_controllers.DeleteTeamController)
}
