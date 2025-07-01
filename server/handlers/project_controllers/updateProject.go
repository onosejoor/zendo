package project_controllers

import (
	"fmt"
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Payload struct {
	Name        string `json:"name" bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}

func UpdateProjectController(ctx *fiber.Ctx) error {
	projectId := ctx.Params("id")
	user := ctx.Locals("user").(*models.UserRes)

	var updatePayload Payload
	if err := ctx.BodyParser(&updatePayload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	collection := db.GetClient().Collection("projects")
	objectId, err := primitive.ObjectIDFromHex(projectId)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid project ID",
		})
	}

	update := bson.M{
		"$set": bson.M{
			"description": updatePayload.Description,
			"name":        updatePayload.Name,
		},
	}

	err = collection.FindOneAndUpdate(
		ctx.Context(),
		bson.M{"_id": objectId, "userId": user.ID},
		update,
	).Err()

	if err != nil {
		log.Println("Error updating project:", err.Error())
		if err == mongo.ErrNoDocuments {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("project with id %v not found", projectId),
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error updating project",
		})
	}

	cacheKey := []string{fmt.Sprintf("user:%s:project:%s", user.ID.Hex(), projectId), fmt.Sprintf("user:%s:projects", user.ID.Hex())}

	if err := redis.DeleteCache(ctx.Context(), cacheKey...); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Project updated",
	})
}
