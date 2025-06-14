package main

import (
	"log"
	"main/db"
	auth_controllers "main/handlers/auth_controllers"
	"main/handlers/task_controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	auth := app.Group("/auth")

	auth.Get("/user/:id", auth_controllers.HandleGetUser)
	auth.Post("/signup", auth_controllers.HandleSignup)
	auth.Post("/signin", auth_controllers.HandleSignin)
	auth.Post("/oauth", auth_controllers.HandleOauth)

	// task
	taskRoute := app.Group("/task")

	taskRoute.Get("/:id", task_controllers.GetTaskByIdController)
	app.Post("/task", task_controllers.CreateTaskController)

	db.GetClient()
	log.Println("Server listening on port 8080")
	app.Listen(":8080")
}
