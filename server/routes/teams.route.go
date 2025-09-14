package routes

import (
	"main/handlers/team_controllers"
	"main/handlers/team_task_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TeamsRoutes(app fiber.Router) {
	teamsRoute := app.Group("/teams")

	teamsRoute.Use(middlewares.AuthMiddleware)

	teamsRoute.Get("/", team_controllers.GetTeamsController)
	teamsRoute.Get("/stats", team_controllers.GetTeamStatsController)
	teamsRoute.Post("/new", team_controllers.CreateTeamController)
	teamsRoute.Get("/teams/members/invite", team_controllers.CreateTeamMemberController)

	teamsRoute.Get("/:teamId", middlewares.RequireTeamMember(), team_controllers.GetTeamByIDController)
	teamsRoute.Put("/:teamId", middlewares.RequireTeamMember("owner"), team_controllers.UpdateTeamController)
	teamsRoute.Get("/:teamId/members", middlewares.RequireTeamMember(), team_controllers.GetTeamMembersController)
	teamsRoute.Post("/:teamId/members/invite", middlewares.RequireTeamMember("admin", "owner"), team_controllers.SendTeamInvite)
	teamsRoute.Delete("/:teamId", middlewares.RequireTeamMember("owner"), team_controllers.DeleteTeamController)

	teamsRoute.Get("/:teamId/tasks", middlewares.RequireTeamMember(), team_task_controllers.GetTeamTasksController)
	teamsRoute.Delete("/:teamId/members/:memberId", middlewares.RequireTeamMember("owner"), team_controllers.RemoveTeamMemberController)
	teamsRoute.Get("/:teamId/tasks/:id", middlewares.RequireTeamMember(), team_task_controllers.GetTeamTaskByIdController)
	teamsRoute.Delete("/:teamId/tasks/:taskId", middlewares.RequireTeamMember("owner"), team_task_controllers.DeleteTeamTaskController)

}
