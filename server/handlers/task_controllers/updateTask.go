package task_controllers

import (
	"context"
	"errors"
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"
	"main/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Payload struct {
	Title       string               `json:"title" bson:"title,omitempty"`
	Description string               `json:"description" bson:"description,omitempty"`
	SubTasks    []models.SubTask     `json:"subTasks,omitempty" bson:"subTasks,omitempty"`
	ProjectId   primitive.ObjectID   `json:"projectId,omitempty" bson:"projectId,omitempty"`
	DueDate     time.Time            `json:"dueDate" bson:"dueDate"`
	Status      string               `json:"status" bson:"status,omitempty"`
	TeamID      primitive.ObjectID   `json:"team_id,omitempty" bson:"team_id,omitempty"`
	Assignees   []primitive.ObjectID `json:"assignees" bson:"assignees"`
}

func UpdateTaskController(ctx *fiber.Ctx) error {
	taskId := utils.HexToObjectID(ctx.Params("id"))
	user := ctx.Locals("user").(*models.UserRes)

	var updatePayload Payload
	if err := ctx.BodyParser(&updatePayload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	client := db.GetClient()
	collection := client.Collection("tasks")

	var task models.Task
	err := collection.FindOne(ctx.Context(), bson.M{"_id": taskId}, findOneOpts).Decode(&task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false,
				"message": "Task not found",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal error",
		})
	}

	if !canModifyTask(&task, user, ctx.Context()) {
		return ctx.Status(403).JSON(fiber.Map{
			"success": false,
			"message": "You are not authorized to update this task",
		})
	}

	update := bson.M{
		"$set": updatePayload,
	}

	err = collection.FindOneAndUpdate(
		ctx.Context(),
		bson.M{"_id": taskId},
		update,
	).Err()

	if err != nil {
		log.Println("Error updating task:", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error updating task",
		})
	}

	if task.TeamID != primitive.NilObjectID {
		redis.ClearTeamMembersCache(ctx.Context(), task.TeamID)
	} else {
		redis.DeleteTaskCache(ctx.Context(), user.ID.Hex())
	}

	if updatePayload.Status != "completed" && updatePayload.DueDate.After(time.Now().Add(10*time.Minute)) {
		go setReminders(taskId, task, context.Background(), user.ID)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task updated",
	})
}
