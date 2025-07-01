package project_controllers

import (
	"log"
	redis "main/configs/redis"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateProjectController(ctx *fiber.Ctx) error {
	var payload models.Project
	user := ctx.Locals("user").(*models.UserRes)

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	if err := utils.Validate.Struct(payload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid fields",
		})
	}

	id, err := models.CreateProject(payload, user.ID, ctx.Context())
	if err != nil {
		log.Println("Error creating project: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "error creating projects, try again",
		})
	}

	if err := redis.DeleteTaskCache(ctx.Context(), user.ID.Hex(), ""); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(201).JSON(fiber.Map{
		"success":   true,
		"projectId": id.Hex(),
	})
}
