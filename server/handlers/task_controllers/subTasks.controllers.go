package task_controllers

import (
	"context"
	"errors"
	"log"
	prometheus "main/configs/prometheus"
	redis "main/configs/redis"
	"main/db"
	"main/models"
	"main/utils"
	"slices"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubTaskPayload struct {
	Completed bool `json:"completed" bson:"completed,omitempty"`
}

var findOneOpts = options.FindOne().SetProjection(bson.M{"userId": 1, "assignees": 1, "team_id": 1})

func canModifyTask(task *models.Task, user *models.UserRes, ctx context.Context) bool {
	if task.UserId == user.ID {
		return true
	}
	if slices.Contains(task.Assignees, user.ID) {
		return true
	}
	if task.TeamID != primitive.NilObjectID {
		return models.CheckMemberRoleMatch(user.ID, task.TeamID, ctx, []string{"owner"})
	}
	return false
}

func clearTaskCache(ctx context.Context, task *models.Task, userID primitive.ObjectID) {
	if task.TeamID != primitive.NilObjectID {
		go redis.ClearTeamMembersCache(context.Background(), task.TeamID)
	} else {
		redis.DeleteTaskCache(ctx, userID.Hex())
		prometheus.RecordRedisOperation("delete_task_cache")
	}
}

func UpdateSubTaskController(ctx *fiber.Ctx) error {
	taskId := utils.HexToObjectID(ctx.Params("id"))
	subTaskId := ctx.Params("subTaskId")

	user := ctx.Locals("user").(*models.UserRes)

	var updatePayload SubTaskPayload
	if err := ctx.BodyParser(&updatePayload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"success": false, "message": err.Error()})
	}

	collection := db.GetClient().Collection("tasks")

	var task models.Task
	if err := collection.FindOne(ctx.Context(),
		bson.M{"_id": taskId},
		findOneOpts,
	).Decode(&task); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(404).JSON(fiber.Map{"success": false, "message": "Task not found"})
		}
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Error finding task"})
	}

	if !canModifyTask(&task, user, ctx.Context()) {
		return ctx.Status(403).JSON(fiber.Map{"success": false, "message": "Not authorized to update this subtask"})
	}

	update := bson.M{"$set": bson.M{"subTasks.$[elem].completed": updatePayload.Completed}}
	arrayFilters := options.ArrayFilters{Filters: bson.A{bson.M{"elem._id": subTaskId}}}

	res, err := collection.UpdateOne(
		ctx.Context(),
		bson.M{"_id": taskId},
		update,
		options.Update().SetArrayFilters(arrayFilters),
	)
	if err != nil {
		log.Println("Error updating subtask:", err)
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Error updating subtask"})
	}
	if res.MatchedCount == 0 {
		return ctx.Status(404).JSON(fiber.Map{"success": false, "message": "Subtask not found"})
	}

	clearTaskCache(ctx.Context(), &task, user.ID)

	return ctx.Status(200).JSON(fiber.Map{"success": true, "message": "SubTask updated"})
}

// 🔹 delete subtask
func DeleteSubTaskController(ctx *fiber.Ctx) error {
	taskId := utils.HexToObjectID(ctx.Params("id"))
	subTaskId := ctx.Params("subTaskId")

	user := ctx.Locals("user").(*models.UserRes)
	collection := db.GetClient().Collection("tasks")

	var task models.Task
	if err := collection.FindOne(ctx.Context(),
		bson.M{"_id": taskId},
		findOneOpts,
	).Decode(&task); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ctx.Status(404).JSON(fiber.Map{"success": false, "message": "Task not found"})
		}
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Error finding task"})
	}

	if !canModifyTask(&task, user, ctx.Context()) {
		return ctx.Status(403).JSON(fiber.Map{"success": false, "message": "Not authorized"})
	}

	update := bson.M{"$pull": bson.M{"subTasks": bson.M{"_id": subTaskId}}}

	res, err := collection.UpdateOne(ctx.Context(), bson.M{"_id": taskId}, update)
	if err != nil {
		log.Println("Error deleting subtask:", err)
		return ctx.Status(500).JSON(fiber.Map{"success": false, "message": "Error deleting subtask"})
	}
	if res.MatchedCount == 0 {
		return ctx.Status(404).JSON(fiber.Map{"success": false, "message": "Subtask not found"})
	}

	clearTaskCache(ctx.Context(), &task, user.ID)

	return ctx.Status(200).JSON(fiber.Map{"success": true, "message": "SubTask deleted"})
}
