package task_controllers

import (
	"log"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func CreateTaskController(ctx *fiber.Ctx) error {
	var body models.Task

	if err := ctx.BodyParser(&body); err != nil {
		log.Println("Error parsing body: ", err)
		return ctx.Status(400).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	id, err := models.CreateTask(body, ctx.Context())
	if err != nil {
		log.Println("Error creating task: ", err)

		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error parsing data",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task created successfully",
		"taskId":  id,
	})
}
