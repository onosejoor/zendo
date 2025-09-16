package team_controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	prometheus "main/configs/prometheus"
	"main/configs/redis"
	"main/db"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Payload struct {
	Name        string `json:"name" bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}

func UpdateTeamController(ctx *fiber.Ctx) error {
	teamId := ctx.Locals("teamId").(primitive.ObjectID)
	user := ctx.Locals("user").(*models.UserRes)

	var updatePayload Payload
	if err := ctx.BodyParser(&updatePayload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	collection := db.GetClient().Collection("teams")

	update := bson.M{
		"$set": bson.M{
			"description": updatePayload.Description,
			"name":        updatePayload.Name,
			"updated_at":  time.Now().Local(),
		},
	}

	err := collection.FindOneAndUpdate(
		ctx.Context(),
		bson.M{"_id": teamId, "creator_id": user.ID},
		update,
	).Err()

	if err != nil {
		log.Println("Error updating team:", err.Error())
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("team with id %v not found", teamId),
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error updating team",
		})
	}

	go redis.ClearTeamMembersCache(context.Background(), teamId)
	prometheus.RecordRedisOperation("delete_cache")
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "team updated",
	})
}
