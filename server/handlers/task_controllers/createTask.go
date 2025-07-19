package task_controllers

import (
	"log"
	"main/configs/cron"
	"main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTaskController(ctx *fiber.Ctx) error {
	var body models.Task

	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error parsing body: ", err)
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	if err := utils.Validate.Struct(body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "All fields must be valid",
		})
	}

	userId := ctx.Locals("user").(*models.UserRes).ID

	id, err := models.CreateTask(body, ctx.Context(), userId)
	if err != nil {
		log.Println("Error creating task: ", err)

		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	cronData := cron.ReminderProps{
		TaskID:   id.(primitive.ObjectID),
		TaskName: body.Title,
		Ctx:      ctx.Context(),
		UserID:   userId,
		DueDate:  body.DueDate.Local(),
	}

	err = cronData.ScheduleReminderJob()
	statusText := "Task created successfully"
	if err != nil {
		statusText = "Task created, but reminder may not be scheduled"
		log.Println("[Scheduler] Failed to schedule reminder: ", err)
	}

	redis.ClearAllCache(ctx.Context(), userId.Hex(), "", body.ProjectId.Hex())

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": statusText,
		"taskId":  id,
	})
}
