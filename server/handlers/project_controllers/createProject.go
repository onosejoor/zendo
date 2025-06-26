package project_controllers

import (
	"fmt"
	"log"
	redis "main/configs"
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

	cacheKey := fmt.Sprintf("user:%s:projects", user.ID.Hex())

	if err := redis.DeleteCache(ctx.Context(), cacheKey); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(201).JSON(fiber.Map{
		"success":   true,
		"projectId": id.Hex(),
	})
}
