package routes

import (
	"main/handlers/team_controllers"
	"main/handlers/team_task_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TeamsRoutes(app fiber.Router) {

	app.Get("/teams/members/invite", team_controllers.CreateTeamMemberController)
	teamsRoute := app.Group("/teams")

	teamsRoute.Use(middlewares.AuthMiddleware)

	teamsRoute.Get("/", team_controllers.GetTeamsController)
	teamsRoute.Post("/", team_controllers.CreateTeamController)
	teamsRoute.Get("/:id", team_controllers.GetTeamByIDController)
	teamsRoute.Get("/:id/members", team_controllers.GetTeamMembersController)
	teamsRoute.Post("/:id/members/invite", team_controllers.SendTeamInvite)
	teamsRoute.Delete("/:id", team_controllers.DeleteTeamController)

	teamsRoute.Get("/:teamId/tasks", team_task_controllers.GetTeamTasksController)
	teamsRoute.Delete("/:teamId/members/:memberId", team_controllers.RemoveTeamMemberController)
	teamsRoute.Get("/:teamId/tasks/:id", team_task_controllers.GetTeamTaskByIdController)
}
