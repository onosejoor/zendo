package auth_controllers

import (
	"log"
	prometheus "main/configs/prometheus"
	"main/cookies"
	"main/db"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleSignup(ctx *fiber.Ctx) error {
	var body models.UserPayload

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
		})
	}

	if err := utils.Validate.Struct(body); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	collection := db.GetClient().Collection("users")
	count, err := collection.CountDocuments(ctx.Context(), bson.M{
		"email": body.Email,
	})
	if err != nil {
		log.Println("DB error:", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Internal server error",
		})
	}

	if count > 0 {
		return ctx.Status(409).JSON(fiber.Map{
			"success": false,
			"message": "User already exists",
		})
	}

	id, err := body.CreateUser(collection, ctx.Context())
	if err != nil {
		log.Println("error creating user: ", err)
		return ctx.Status(500).JSON(fiber.Map{
			"success": false, "message": "Internal error",
		})
	}
	prometheus.RecordUserRegistration()
	user := models.UserRes{Username: body.Username, ID: id, EmailVerified: false}

	token, err := cookies.GenerateEmailToken(&user)
	if err != nil {
		log.Println("Error Generating email token: ", err.Error())
		return ctx.Status(201).JSON(fiber.Map{
			"success":    true,
			"message":    "User created successfully, but email token not sent",
			"email_sent": false,
		})
	}

	err = cookies.SendEmail(token, models.UserPayload{
		Username: body.Username,
		Email:    body.Email,
	})
	if err != nil {
		log.Println("Error sending email: ", err.Error())
		return ctx.Status(201).JSON(fiber.Map{
			"success":    true,
			"message":    "User created successfully, but email not verified. Re-login to continue",
			"email_sent": false,
		})
	}

	// This code was commented out to not allow new users access route except email is verified
	// err = cookies.CreateSession(user, ctx)
	// if err != nil {
	// 	return ctx.Status(500).JSON(fiber.Map{
	// 		"success":    true,
	// 		"message":    "Internal error",
	// 		"email_sent": false,
	// 	})
	// }

	return ctx.Status(201).JSON(fiber.Map{
		"success":    true,
		"message":    "Verification Link Sent To Email",
		"email_sent": true,
	})

}
