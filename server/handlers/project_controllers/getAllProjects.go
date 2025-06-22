package project_controllers

import (
	"log"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllProjectsController(ctx *fiber.Ctx) error {

	client := db.GetClient()
	collection := client.Collection("projects")

	user := ctx.Locals("user").(*models.UserRes)

	cursor, err := collection.Find(ctx.Context(), bson.M{"ownerId": user.ID})
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	var projects = make([]models.Project, 0)

	if err := cursor.All(ctx.Context(), &projects); err != nil {
		log.Println("Error parsing db data to slice: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	return ctx.Status(200).JSON(bson.M{
		"success":  true,
		"projects": projects,
	})

}
