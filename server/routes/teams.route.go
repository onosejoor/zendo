package routes

import (
	"main/handlers/team_controllers"
	"main/handlers/team_task_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TeamsRoutes(app fiber.Router) {
	teams := app.Group("/teams", middlewares.AuthMiddleware)

	// top-level
	teams.Get("/", team_controllers.GetTeamsController)
	teams.Get("/stats", team_controllers.GetTeamStatsController)
	teams.Post("/new", team_controllers.CreateTeamController)
	teams.Get("/members/invite", team_controllers.CreateTeamMemberController)

	// team-specific
	team := teams.Group("/:teamId", middlewares.RequireTeamMember())

	team.Get("/", team_controllers.GetTeamByIDController)
	team.Get("/stats", team_controllers.GetTeamByIdStatsController)
	team.Get("/invites", middlewares.RequireTeamMember("owner"), team_controllers.GetTeamInvitesController)
	team.Put("/", middlewares.RequireTeamMember("owner"), team_controllers.UpdateTeamController)
	team.Delete("/invites/:inviteId", middlewares.RequireTeamMember("owner"), team_controllers.CancelInviteController)
	team.Delete("/", middlewares.RequireTeamMember("owner"), team_controllers.DeleteTeamController)

	// members
	members := team.Group("/members")
	members.Get("/", team_controllers.GetTeamMembersController)
	members.Post("/invite", middlewares.RequireTeamMember("admin", "owner"), team_controllers.SendTeamInvite)
	members.Delete("/:memberId", team_controllers.RemoveTeamMemberController)

	// tasks
	tasks := team.Group("/tasks")
	tasks.Get("/", team_task_controllers.GetTeamTasksController)
	tasks.Get("/:id", team_task_controllers.GetTeamTaskByIdController)
	tasks.Delete("/:taskId", middlewares.RequireTeamMember("owner"), team_task_controllers.DeleteTeamTaskController)
}
