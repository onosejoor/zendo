package task_controllers

import (
	"fmt"
	"log"
	redis "main/configs/redis"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubTaskPayload struct {
	Completed bool `json:"completed" bson:"completed,omitempty"`
}

func UpdateSubTaskController(ctx *fiber.Ctx) error {
	taskId := ctx.Params("id")
	subTaskIdParam := ctx.Params("subTaskId")

	user := ctx.Locals("user").(*models.UserRes)

	var updatePayload SubTaskPayload
	if err := ctx.BodyParser(&updatePayload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	collection := db.GetClient().Collection("tasks")
	objectId, _ := primitive.ObjectIDFromHex(taskId)

	update := bson.M{
		"$set": bson.M{
			"subTasks.$[elem].completed": updatePayload.Completed,
		},
	}

	arrayFilters := options.ArrayFilters{
		Filters: bson.A{
			bson.M{"elem._id": subTaskIdParam},
		},
	}

	_, err := collection.UpdateOne(
		ctx.Context(),
		bson.M{"_id": objectId, "userId": user.ID},
		update,
		options.Update().SetArrayFilters(arrayFilters),
	)

	if err != nil {
		log.Println("Error updating subtask:", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error updating subtask",
		})
	}

	if err := redis.DeleteTaskCache(ctx.Context(), user.ID.Hex(), taskId); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "SubTask updated",
	})
}

func DeleteSubTaskController(ctx *fiber.Ctx) error {
	taskId := ctx.Params("id")
	subTaskIdParam := ctx.Params("subTaskId")

	user := ctx.Locals("user").(*models.UserRes)

	collection := db.GetClient().Collection("tasks")
	objectId, _ := primitive.ObjectIDFromHex(taskId)

	update := bson.M{
		"$pull": bson.M{
			"subTasks": bson.M{
				"_id": subTaskIdParam,
			},
		},
	}

	_, err := collection.UpdateOne(
		ctx.Context(),
		bson.M{"_id": objectId, "userId": user.ID},
		update,
	)
	if err != nil {
		log.Println("Error deleting sub_task:", err.Error())
		if err == mongo.ErrNoDocuments {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("Subtask with id %v not found", subTaskIdParam),
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error deleting subtask",
		})
	}

	if err := redis.DeleteTaskCache(ctx.Context(), user.ID.Hex(), taskId); err != nil {
		log.Println(err.Error())
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "SubTask deleted",
	})
}
