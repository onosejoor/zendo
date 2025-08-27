package task_controllers

import (
	"log"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSearchedTasksController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	search := ctx.Query("search")
	var dbTaks = make([]models.Task, 0)

	client := db.GetClient()
	collection := client.Collection("tasks")

	filter := bson.M{"userId": user.ID}

	if search != "" {
		filter["$or"] = []bson.M{
			{"title": bson.M{"$regex": search, "$options": "i"}},
			{"description": bson.M{"$regex": search, "$options": "i"}},
		}
	}

	cursor, err := collection.Find(ctx.Context(), filter)
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

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
