package main

import (
	"log"
	"main/db"
	"main/handlers/auth_controllers"
	"main/handlers/project_controllers"
	"main/handlers/task_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env: ", err.Error())
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	// auth
	auth := app.Group("/auth")

	auth.Get("/user", middlewares.AuthMiddleware, auth_controllers.HandleGetUser)
	auth.Post("/signup", auth_controllers.HandleSignup)
	auth.Post("/signin", auth_controllers.HandleSignin)
	auth.Post("/oauth", auth_controllers.HandleOauth)

	// task
	taskRoute := app.Group("/task")
	taskRoute.Use(middlewares.AuthMiddleware)

	taskRoute.Get("", task_controllers.GetAllTasksController)
	taskRoute.Get("/:id", task_controllers.GetTaskByIdController)
	taskRoute.Post("/new", task_controllers.CreateTaskController)
	taskRoute.Put("/:id", task_controllers.UpdateTaskController)
	taskRoute.Delete("/:id", task_controllers.DeleteTaskController)

	// projects
	projectRoute := app.Group("/projects")
	projectRoute.Use(middlewares.AuthMiddleware)

	projectRoute.Get("", project_controllers.GetAllProjectsController)
	projectRoute.Get("/:id/tasks", project_controllers.GetTaskInProjectsController)
	projectRoute.Get("/:id", project_controllers.GetProjectByIdController)
	projectRoute.Post("/new", project_controllers.CreateProjectController)
	projectRoute.Delete("/:id", project_controllers.DeleteProjectController)

	db.GetClient()
	log.Println("Server listening on port 8080")
	app.Listen(":8080")
}
