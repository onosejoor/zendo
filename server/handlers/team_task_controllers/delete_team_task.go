package team_task_controllers

import (
	"context"
	"fmt"
	"log"
	"main/configs/prometheus"
	"main/configs/redis"
	"main/db"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTeamTaskController(ctx *fiber.Ctx) error {
	role := ctx.Locals("role").(string)

	teamID := ctx.Locals("teamId").(primitive.ObjectID)
	taskID := utils.HexToObjectID(ctx.Params("taskId"))

	if role != "owner" {
		return ctx.Status(403).JSON(fiber.Map{
			"success": false,
			"message": "Only team owners can delete tasks",
		})
	}

	collection := db.GetClient().Collection("tasks")

	_, err := collection.DeleteOne(ctx.Context(), bson.M{"_id": taskID, "team_id": teamID})
	if err != nil {
		log.Println("Error deleting task:", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to delete task",
		})
	}

	go redis.ClearTeamMembersCache(context.Background(), teamID)
	prometheus.RecordRedisOperation("delete_team_task_cache")

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": fmt.Sprintf("Task %s deleted successfully", taskID.Hex()),
	})
}
