package task_controllers

import (
	"fmt"
	"log"
	"main/configs/redis"
	"main/models"
	"main/utils"

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

	cacheKey := []string{fmt.Sprintf("user:%s:tasks", userId.Hex())}

	if err := redis.DeleteCache(ctx.Context(), cacheKey...); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Task created successfully",
		"taskId":  id,
	})
}
