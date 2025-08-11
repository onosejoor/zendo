package task_controllers

import (
	"log"
	"main/configs/cron"
	"main/configs/redis"
	"main/models"
	"main/utils"
	"time"

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

	if !body.DueDate.After(time.Now().Local()) {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Due date must be in the future",
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

	reminderPayload := models.Reminder{
		TaskID:     id.(primitive.ObjectID),
		TaskName:   body.Title,
		UserID:     userId,
		DueDate:    body.DueDate.Local(),
		Expires_At: body.DueDate.Local(),
		CreatedAt:  time.Now().Local(),
	}

	statusText := "Task created successfully"

	err = reminderPayload.CreateReminder(ctx.Context())
	if err != nil {
		statusText = "Task created, but reminder may not be scheduled"
		log.Println("CREATE REMINDER FAILED: ", err)
	}

	if !body.DueDate.After(time.Now().Local().Add(10*time.Minute)) && body.Status != "completed" {
		payload := cron.ReminderProps{
			TaskID:   id.(primitive.ObjectID),
			TaskName: body.Title,
			UserID:   userId,
			DueDate:  body.DueDate.Local(),
			Ctx:      ctx.Context(),
		}

		err := payload.ScheduleReminderJob()
		if err != nil {
			statusText = "Task created, but reminder may not be scheduled"
			log.Println("[Scheduler] Failed to schedule reminder: ", err)
		}

	}

	redis.ClearAllCache(ctx.Context(), userId.Hex(), "", body.ProjectId.Hex())

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": statusText,
		"taskId":  id,
	})
}
