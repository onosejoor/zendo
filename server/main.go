package main

import (
	"log"
	"main/auth/handlers"
	"main/db"

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

	auth.Get("/user/:id", handlers.HandleGetUser)
	auth.Post("/signup", handlers.HandleSignup)
	auth.Post("/signin", handlers.HandleSignin)
	auth.Post("/oauth", handlers.HandleOauth)

	db.GetClient()
	log.Println("Server listening on port 8080")
	app.Listen(":8080")
}
