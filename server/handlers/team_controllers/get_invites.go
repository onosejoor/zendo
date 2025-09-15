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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TeamInviteRes struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email" bson:"email"`
	Status    string             `json:"status" bson:"status"`
	ExpiresAt time.Time          `json:"expiresAt" bson:"expiresAt"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

func GetTeamInvitesController(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*models.UserRes)
	teamId := ctx.Locals("teamId").(primitive.ObjectID)

	cacheKey := fmt.Sprintf("user:%s:teams:%s:invitees", user.ID.Hex(), teamId.Hex())

	var invitees = make([]TeamInviteRes, 0)

	redisClient := redis.GetRedisClient()
	if redisClient.GetCacheHandler(ctx, &invitees, cacheKey, "invitees") {
		return nil
	}

	collection := db.GetClient().Collection("team_invites")

	opts := options.Find().SetProjection(bson.M{"email": 1, "expiresAt": 1, "status": 1, "createdAt": 1})

	cursor, err := collection.Find(ctx.Context(), bson.M{"team_id": teamId}, opts)
	if err != nil {
		log.Println("Error getting pending users: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": true, "message": "Error fetching pending users, try again",
		})
	}
	defer cursor.Close(ctx.Context())

	if err := cursor.All(ctx.Context(), &invitees); err != nil {
		log.Println(err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": true, "message": "Internal server error",
		})
	}

	redisClient.SetCacheData(cacheKey, ctx.Context(), invitees)

	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "invitees": invitees,
	})
}
