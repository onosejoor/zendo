package task_controllers

import (
	"log"
	"main/configs/cron"
	redis "main/configs/redis"
	"main/db"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteTaskController(ctx *fiber.Ctx) error {
	taskId := ctx.Params("id")

	user := ctx.Locals("user").(*models.UserRes)

	objectId, _ := primitive.ObjectIDFromHex(taskId)

	client := db.GetClient()
	collection := client.Collection("tasks")

	var task models.Task

	err := collection.FindOne(ctx.Context(), bson.M{"_id": objectId, "userId": user.ID}).Decode(&task)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "Task not found",
			})

		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error, try again",
		})
	}

	if task.ProjectId != primitive.NilObjectID {
		err := task.DeleteTaskWithTransaction(ctx.Context())
		if err != nil {
			log.Println("Error deleting task: ", err)

			return ctx.Status(500).JSON(fiber.Map{
				"success": false, "message": "Error Deleting Task",
			})
		}
	} else {
		if _, err := collection.DeleteOne(ctx.Context(), bson.M{"_id": objectId, "userId": user.ID}); err != nil {
			log.Println("Error deleting task: ", err.Error())
			return ctx.Status(500).JSON(fiber.Map{
				"success": false, "message": "Internal error",
			})
		}
	}

	err = models.DeleteReminder(ctx.Context(), objectId, user.ID)
	if err != nil {
		log.Println("Error deleting reminder: ", err)
		err = models.DeleteReminder(ctx.Context(), objectId, user.ID)
		if err != nil {
			log.Println("Error deleting reminder (RETRY): ", err)
		}
	}
	if !time.Now().After(task.DueDate) && !task.DueDate.After(time.Now().Local().Add(10*time.Minute)) {
		cron.DeleteCronJob(objectId)
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex(), task.ID.Hex(), task.ProjectId.Hex())

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Task deleted",
	})
}

func DeleteAllTasksController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)

	err := models.DeleteAllTasksWithTransaction(ctx.Context(), user)
	if err != nil {
		log.Println("Error deleing all tasks: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error deleting all tasks, try again",
		})
	}

	err = models.DeleteAllReminder(ctx.Context(), user.ID)
	if err != nil {
		log.Println("Error deleting all reminders: ", err)
		err := models.DeleteAllReminder(ctx.Context(), user.ID)
		if err != nil {
			log.Println("Error deleting all reminders (RETRY): ", err)
		}
	}

	redis.ClearAllCache(ctx.Context(), user.ID.Hex(), "", "")

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "Task deleted",
	})
}
