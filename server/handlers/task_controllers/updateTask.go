package task_controllers

import (
	"fmt"
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Payload struct {
	Title       string             `json:"title" bson:"title,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	UserId      primitive.ObjectID `json:"userId" bson:"userId"`
	SubTasks    []models.SubTasks  `json:"subTasks,omitempty" bson:"subTasks,omitempty"`
	ProjectId   primitive.ObjectID `json:"projectId,omitempty" bson:"projectId,omitempty"`
	Status      string             `json:"status" bson:"status,omitempty"`
}

func UpdateTaskController(ctx *fiber.Ctx) error {
	taskId := ctx.Params("id")
	user := ctx.Locals("user").(*models.UserRes)

	var updatePayload Payload
	if err := ctx.BodyParser(&updatePayload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	collection := db.GetClient().Collection("tasks")
	objectId, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid task ID",
		})
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updatePayload.Title,
			"description": updatePayload.Description,
			"subTasks":    updatePayload.SubTasks,
			"projectId":   updatePayload.ProjectId,
			"status":      updatePayload.Status,
		},
	}

	err = collection.FindOneAndUpdate(
		ctx.Context(),
		bson.M{"_id": objectId, "userId": user.ID},
		update,
	).Err()

	if err != nil {
		log.Println("Error updating task:", err.Error())
		if err == mongo.ErrNoDocuments {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("Task with id %v not found", taskId),
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error updating task",
		})
	}

	cacheKey := []string{fmt.Sprintf("user:%s:task:%s", user.ID.Hex(), taskId), fmt.Sprintf("user:%s:tasks", user.ID.Hex())}

	if err := redis.DeleteCache(ctx.Context(), cacheKey...); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task updated",
	})
}
