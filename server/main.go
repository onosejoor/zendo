package main

import (
	"context"
	"log"
	"main/configs/cron"
	prometheus_config "main/configs/prometheus"
	redis "main/configs/redis"
	"main/db"
	"main/handlers"
	"main/middlewares"
	"main/routes"
	"main/utils"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	utils.PullEnv()

	prometheus_config.Init()
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

	app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${method} | ${path} | ${latency}\n",
	}))

	app.Use(prometheus_config.NewMiddleware())

	app.Get("/metrics", handlers.MetricsHandler)
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"success": true, "message": "Healthy",
		})
	})

	// stats
	app.Get("/stats", middlewares.AuthMiddleware, handlers.GetStatsControllers)

	// auth
	routes.AuthRoutes(app)

	// task
	routes.TaskRoutes(app)

	// projects
	routes.ProjectRoutes(app)

	// Teams and Team Members
	routes.TeamsRoutes(app)

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
