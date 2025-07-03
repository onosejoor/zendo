package handlers

import (
	"log"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type Stats struct {
	TotalTasks     int `json:"total_tasks"`
	TotalProjects  int `json:"total_projects"`
	CompletionRate int `json:"completion_rate"`
	CompletedTasks int `json:"completed_tasks"`
	DueToday       int `json:"dueToday"`
}

func GetStatsControllers(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)

	var dbTaks = make([]models.Task, 0)

	var stats Stats

	client := db.GetClient()

	cursor, err := client.Collection("tasks").Find(ctx.Context(), bson.M{
		"userId": user.ID,
	})
	if err != nil {
		log.Println("Error getting stats: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "error getting stats",
		})
	}

	if err := cursor.All(ctx.Context(), &dbTaks); err != nil {
		log.Println("Error parsing db data to slice: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	noOfProjects, err := client.Collection("projects").CountDocuments(ctx.Context(), bson.M{
		"userId": user.ID,
	})
	if err != nil {
		log.Println("Error getting stats: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "error getting stats",
		})
	}

	noOfTasksDueToday, completionRate, completedTasks := models.GetCompletionRateAndDueDate(dbTaks)

	stats = Stats{
		TotalTasks:     len(dbTaks),
		DueToday:       noOfTasksDueToday,
		CompletionRate: completionRate,
		TotalProjects:  int(noOfProjects),
		CompletedTasks: completedTasks,
	}

	// _ = redisClient.SetCacheData(cacheKey, ctx.Context(), stats)

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"stats":   stats,
	})

}
