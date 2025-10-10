package handlers

import (
	"log"
	"main/configs/prometheus"
	"main/configs/redis"
	"main/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Avatars    []string `json:"avatars"`
	TotalUsers int64    `json:"total_users"`
}

type Avatars struct {
	Avatar string `bson:"avatar"`
}

func GetHomePageData(ctx *fiber.Ctx) error {

	var data Response

	client := db.GetClient()

	cacheKey := "user:home_stats"

	redisClient := redis.GetRedisClient()

	if redisClient.GetCacheHandler(ctx, &data, cacheKey, "data") {
		prometheus.RecordRedisOperation("get_cache")
		return nil
	}
	prometheus.RecordRedisOperation("cache_miss")

	totalUsers, err := client.Collection("users").CountDocuments(ctx.Context(), bson.M{})
	if err != nil {
		log.Println("Error getting response: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "error getting data",
		})
	}

	opts := options.Find().SetLimit(5).SetProjection(bson.M{"_id": 0, "avatar": 1})
	cursor, err := client.Collection("users").Find(ctx.Context(), bson.M{"avatar": bson.M{"$ne": ""}}, opts)
	if err != nil {
		log.Println("Error getting response: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "error getting data",
		})
	}

	var avatars []struct{ Avatar string }
	if err := cursor.All(ctx.Context(), &avatars); err != nil {
		log.Println("Error getting data: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "error getting data",
		})
	}

	avatarStrings := make([]string, len(avatars))
	for i, v := range avatars {
		avatarStrings[i] = v.Avatar
	}

	data = Response{
		TotalUsers: totalUsers,
		Avatars:    avatarStrings,
	}

	_ = redisClient.SetCacheData(cacheKey, ctx.Context(), data)
	prometheus.RecordRedisOperation("set_cache")
	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})

}
