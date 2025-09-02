package routes

import (
	"main/handlers/project_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(app fiber.Router) {
	projectRoute := app.Group("/projects")
	projectRoute.Use(middlewares.AuthMiddleware)

	projectRoute.Get("", project_controllers.GetAllProjectsController)
	projectRoute.Get("/search", project_controllers.GetSearchedProjectsController)
	projectRoute.Get("/:id/tasks", project_controllers.GetTaskInProjectsController)
	projectRoute.Get("/:id", project_controllers.GetProjectByIdController)
	projectRoute.Post("/new", project_controllers.CreateProjectController)
	projectRoute.Put("/:id", project_controllers.UpdateProjectController)
	projectRoute.Delete("/all", project_controllers.DeleteAllProjectsController)
	projectRoute.Delete("/:id", project_controllers.DeleteProjectController)
}
