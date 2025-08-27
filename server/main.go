package main

import (
	"context"
	"log"
	"main/configs/cron"
	oauth_config "main/configs/oauth"
	redis "main/configs/redis"
	"main/db"
	"main/handlers"
	"main/handlers/auth_controllers"
	"main/handlers/email_controllers"
	"main/handlers/project_controllers"
	"main/handlers/task_controllers"
	"main/middlewares"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CLIENT_URL"),
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
	}))

	app.Use(middlewares.LoggerMiddleware)

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"success": true, "message": "Healthy",
		})
	})

	oauth := oauth_config.InitializeOauthConfig()

	// stats
	app.Get("/stats", middlewares.AuthMiddleware, handlers.GetStatsControllers)

	// auth
	auth := app.Group("/auth")

	auth.Get("/user", middlewares.AuthMiddleware, auth_controllers.HandleGetUser)
	auth.Put("/user", middlewares.AuthMiddleware, auth_controllers.UpdateUserController)
	auth.Get("/verify_email", email_controllers.HandleVerifyEmailController)
	auth.Post("/verify_email", middlewares.AuthMiddleware, email_controllers.SendEmailTokenController)
	auth.Get("/refresh-token", auth_controllers.HandleAccessToken)
	auth.Post("/signup", auth_controllers.HandleSignup)
	auth.Post("/signin", auth_controllers.HandleSignin)

	auth.Get("/oauth/google", oauth.GetOauthController)
	auth.Get("/oauth/callback", oauth.OauthCallBackController)
	auth.Post("/oauth/exchange", oauth.OauthExchangeController)

	// task
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

	// projects
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

	cron.InitializeGoCron()
	redis.GetRedisClient()
	client := db.GetClient()
	db.CreateReminderTTLIndex(client.Collection("reminders"))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for {
			select {
			case <-time.After(10 * time.Minute):
				_, err := http.Get(os.Getenv("ORIGIN") + "/health")
				if err != nil {
					log.Printf("failed to fetch health status: %v \n", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	log.Println("Server listening on port: ", port)
	app.Listen(":" + port)
}
