package email_controllers

import (
	"context"
	"log"
	"main/configs/redis"
	"main/cookies"
	"main/db"
	"main/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func HandleVerifyEmailController(ctx *fiber.Ctx) error {
	token := ctx.Query("token")

	claims, err := cookies.VerifyEmailToken(token)
	if err != nil {
		log.Println("Error Verifying email token: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Invalid or expired token",
		})
	}

	objectId, _ := primitive.ObjectIDFromHex(claims.ID)
	err = UpdateUser(objectId, ctx.Context())
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return ctx.Status(404).JSON(fiber.Map{
				"success": false, "message": "User already verified or not found",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	user := models.UserRes{Username: claims.UserName, ID: objectId, EmailVerified: true}

	err = cookies.CreateSession(user, ctx)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal error",
		})
	}

	_ = redis.DeleteUserCache(ctx.Context(), user.ID.Hex())
	return ctx.Status(200).JSON(fiber.Map{
		"success": true, "message": "User verified successfully",
	})

}

func UpdateUser(id primitive.ObjectID, ctx context.Context) error {
	collection := db.GetClient().Collection("users")

	filter := bson.M{
		"_id": id,
		"$or": []bson.M{
			{"email_verified": false},
			{"email_verified": bson.M{"$exists": false}},
		},
	}

	update := bson.M{
		"$set": bson.M{"email_verified": true},
	}

	err := collection.FindOneAndUpdate(ctx, filter, update).Err()
	if err != nil {
		return err
	}
	return nil
}
