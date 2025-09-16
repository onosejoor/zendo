package auth_controllers

import (
	"context"
	"errors"
	"log"
	prometheus "main/configs/prometheus"
	"main/configs/redis"
	"main/cookies"
	"main/db"
	"main/models"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type User struct {
	Username string `bson:"username,omitempty" json:"username,omitempty"`
	Avatar   string `bson:"avatar,omitempty" json:"avatar,omitempty"`
}

func UpdateUserController(ctx *fiber.Ctx) error {
	var payload User
	var user = ctx.Locals("user").(*models.UserRes)

	payload.Username = ctx.FormValue("username")
	file, _ := ctx.FormFile("avatarFile")

	if file == nil {
		payload.Avatar = ctx.FormValue("avatar")
	}

	if err := uploadFile(ctx.Context(), file, &payload); err != nil {
		log.Println("Error uploading image: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	update := bson.M{
		"$set": payload,
	}

	collection := db.GetClient().Collection("users")
	_, err := collection.UpdateOne(ctx.Context(), bson.M{"_id": user.ID}, update)
	if err != nil {
		log.Println("Error updating user: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error updating user",
		})
	}

	_ = redis.DeleteUserCache(ctx.Context(), user.ID.Hex())
	cookies.CreateSession(models.UserRes{Username: payload.Username, ID: user.ID, EmailVerified: user.EmailVerified}, ctx)

	prometheus.RecordRedisOperation("delete_user_cache")
	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "user updated succesfully",
	})

}

func uploadFile(ctx context.Context, fileHeader *multipart.FileHeader, user *User) error {
	if fileHeader == nil {
		return nil
	}

	if fileHeader.Size > 5*1024*1024 {
		return errors.New("file too large (max 5MB)")
	}

	if !strings.HasPrefix(fileHeader.Header.Get("Content-Type"), "image/") {
		return errors.New("invalid file type")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return errors.New("failed to open file")
	}
	defer file.Close()

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		UploadPreset: os.Getenv("UPLOAD_PRESET"),
	})
	if err != nil {
		return errors.New("cloudinary upload failed")
	}

	user.Avatar = resp.SecureURL

	return nil
}
