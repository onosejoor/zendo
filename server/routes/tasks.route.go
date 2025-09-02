package routes

import (
	"main/handlers/task_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app fiber.Router) {
	taskRoute := app.Group("/tasks")
	taskRoute.Use(middlewares.AuthMiddleware)

	taskRoute.Get("", task_controllers.GetAllTasksController)
	taskRoute.Get("/search", task_controllers.GetSearchedTasksController)
	taskRoute.Get("/:id", task_controllers.GetTaskByIdController)
	taskRoute.Post("/new", task_controllers.CreateTaskController)
	taskRoute.Put("/:id", task_controllers.UpdateTaskController)
	taskRoute.Put("/:id/subtask/:subTaskId", task_controllers.UpdateSubTaskController)
	taskRoute.Delete("/:id/subtask/:subTaskId", task_controllers.DeleteSubTaskController)
	taskRoute.Delete("/all", task_controllers.DeleteAllTasksController)
	taskRoute.Delete("/:id", task_controllers.DeleteTaskController)
}
