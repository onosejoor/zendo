package task_controllers

import (
	"fmt"
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Payload struct {
	Title        string             `json:"title" bson:"title,omitempty"`
	Description  string             `json:"description" bson:"description,omitempty"`
	UserId       primitive.ObjectID `json:"userId" bson:"userId"`
	SubTasks     []models.SubTask   `json:"subTasks,omitempty" bson:"subTasks,omitempty"`
	ProjectId    primitive.ObjectID `json:"projectId,omitempty" bson:"projectId,omitempty"`
	DueDate      time.Time          `json:"dueDate" bson:"dueDate"`
	ReminderSent bool               `json:"-" bson:"reminder_sent"`
	Status       string             `json:"status" bson:"status,omitempty"`
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
		"$set": updatePayload,
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

	if err := redis.DeleteTaskCache(ctx.Context(), user.ID.Hex(), taskId); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task updated",
	})
}
