package project_controllers

import (
	"log"
	"main/db"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSearchedProjectsController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var projects = make([]models.Project, 0)

	search := ctx.Query("search")
	page := ctx.Query("page")
	limit := ctx.Query("limit")

	client := db.GetClient()
	collection := client.Collection("projects")

	filter := bson.M{"ownerId": user.ID}

	if search != "" {
		filter["$or"] = []bson.M{
			{"title": bson.M{"$regex": search, "$options": "i"}},
			{"description": bson.M{"$regex": search, "$options": "i"}},
		}
	}
	opts := utils.GeneratePaginationOptions(page, limit)

	cursor, err := collection.Find(ctx.Context(), filter, opts)
	if err != nil {
		log.Println("Error querying db: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

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
