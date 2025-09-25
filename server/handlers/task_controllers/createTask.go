package task_controllers

import (
	"context"
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

	if err := utils.ParseBodyAndValidateStruct(&body, ctx); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	if !body.DueDate.After(time.Now().Local()) {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Due date must be in the future",
		})
	}

	userId := ctx.Locals("user").(*models.UserRes).ID

	if body.TeamID != primitive.NilObjectID {
		isMember, role, err := models.IsTeamMembers(ctx.Context(), []primitive.ObjectID{userId}, body.TeamID, true)
		if err != nil {
			log.Println("Error checking team membership: ", err)
			return ctx.Status(500).JSON(fiber.Map{
				"success": false, "message": "Internal server error",
			})
		}
		if !isMember {
			return ctx.Status(403).JSON(fiber.Map{
				"success": false, "message": "You are not a member of this team",
			})
		}
		if role != "admin" && role != "owner" {
			return ctx.Status(403).JSON(fiber.Map{
				"success": false, "message": "You cannot create a task in this team",
			})
		}

		if len(body.Assignees) > 0 {
			assignees, err := models.CheckAssignee(body.Assignees, body.TeamID, ctx.Context())
			if err != nil {
				log.Println("Error checking assignees: ", err)
				return ctx.Status(500).JSON(fiber.Map{
					"success": false, "message": "Internal server error",
				})
			}

			if !assignees {
				return ctx.Status(400).JSON(fiber.Map{
					"success": false, "message": "One or more assignees are not members of the team",
				})
			}
		}
	}

	id, err := models.CreateTask(body, ctx.Context(), userId)
	if err != nil {
		log.Println("Error creating task: ", err)

		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	go setReminders(id.(primitive.ObjectID), body, context.Background(), userId)

	if len(body.Assignees) > 0 {
		for _, assigneeId := range body.Assignees {
			if assigneeId != userId {
				go setReminders(id.(primitive.ObjectID), body, ctx.Context(), assigneeId)
			}
		}
	}

	redis.ClearAllCache(ctx.Context(), userId.Hex())

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task created successfully",
		"taskId":  id,
	})
}

func setReminders(id primitive.ObjectID, body models.Task, ctx context.Context, userId primitive.ObjectID) {

	reminderPayload := models.Reminder{
		TaskID:     id,
		TaskName:   body.Title,
		UserID:     userId,
		DueDate:    body.DueDate.Local(),
		Expires_At: body.DueDate.Local(),
		CreatedAt:  time.Now().Local(),
	}

	err := reminderPayload.CreateReminder(ctx)
	if err != nil {
		log.Println("CREATE REMINDER FAILED: ", err)
	}

	if !body.DueDate.After(time.Now().Local().Add(10*time.Minute)) && body.Status != "completed" {
		payload := cron.ReminderProps{
			TaskID:   id,
			TaskName: body.Title,
			UserID:   userId,
			DueDate:  body.DueDate.Local(),
			Ctx:      ctx,
		}

		err := payload.ScheduleReminderJob()
		if err != nil {
			log.Println("[Scheduler] Failed to schedule reminder: ", err)
		}
	}
}
