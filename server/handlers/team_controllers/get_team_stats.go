package team_controllers

import (
	"fmt"
	"log"
	"main/configs/redis"
	"main/db"
	"main/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/errgroup"
)

type Response struct {
	NumberOfTeams             int `json:"number_of_teams"`
	NumberOfTasksAssignedToMe int `json:"number_of_tasks_assigned_to_me"`
	NumberOfTasksDueToday     int `json:"number_of_tasks_due_today"`
}

func GetTeamStatsController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	var res Response

	cacheKey := fmt.Sprintf("user:%s:teams:stats", user.ID.Hex())

	redisClient := redis.GetRedisClient()
	if isReturned := redisClient.GetCacheHandler(ctx, &res, cacheKey, "stat"); isReturned {
		return nil
	}

	g, c := errgroup.WithContext(ctx.Context())

	client := db.GetClient()
	teamMembersCollection := client.Collection("team_members")
	taskCollection := client.Collection("tasks")

	g.Go(func() error {
		numberOfTeams, err := teamMembersCollection.CountDocuments(c, bson.M{"user_id": user.ID})
		if err != nil {
			return err
		}
		res.NumberOfTeams = int(numberOfTeams)
		return err
	})

	g.Go(func() error {
		numberOfTasks, err := taskCollection.CountDocuments(c, bson.M{"assignees": bson.M{"$in": []primitive.ObjectID{user.ID}}})
		if err != nil {
			return err
		}
		res.NumberOfTasksAssignedToMe = int(numberOfTasks)
		return err
	})

	g.Go(func() error {
		cursor, err := taskCollection.Aggregate(c, getPipeline(user.ID))
		if err != nil {
			return err
		}
		var results []struct {
			Count int `bson:"count"`
		}
		if err := cursor.All(c, &results); err != nil {
			return err
		}
		if len(results) > 0 {
			res.NumberOfTasksDueToday = results[0].Count
		} else {
			res.NumberOfTasksDueToday = 0
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Println("ERROR LOADING STATS: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), res)
	return ctx.JSON(fiber.Map{
		"success": true, "stat": res,
	})
}

func getPipeline(userId primitive.ObjectID) mongo.Pipeline {
	today := time.Now()
	todayStart := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
	todayEnd := todayStart.Add(24 * time.Hour)

	return mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "team_id", Value: bson.D{{Key: "$ne", Value: nil}}},
			{Key: "dueDate", Value: bson.D{
				{Key: "$gte", Value: todayStart},
				{Key: "$lt", Value: todayEnd},
			}},
		}}},
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "team_members"},
			{Key: "localField", Value: "team_id"},
			{Key: "foreignField", Value: "team_id"},
			{Key: "as", Value: "team_member"},
		}}},
		{{Key: "$unwind", Value: "$team_member"}},
		{{Key: "$match", Value: bson.D{
			{Key: "team_member.user_id", Value: userId},
		}}},
		{{Key: "$count", Value: "count"}},
	}
}
