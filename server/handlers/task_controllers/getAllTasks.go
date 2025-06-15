package task_controllers

import (
	"log"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllTasksController(ctx *fiber.Ctx) error {

	client := db.GetClient()
	collection := client.Collection("tasks")

	user := ctx.Locals("user").(*models.UserRes)

	cursor, err := collection.Find(ctx.Context(), bson.M{"userId": user.ID})
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	var dbTaks []models.Task

	if err := cursor.All(ctx.Context(), &dbTaks); err != nil {
		log.Println("Error parsing db data to slice: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	return ctx.Status(200).JSON(bson.M{
		"success": true,
		"tasks":   dbTaks,
	})

}
