package email_controllers

import (
	"log"
	"main/cookies"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func SendEmailTokenController(ctx *fiber.Ctx) error {

	user := ctx.Locals("user").(*models.UserRes)

	if user.EmailVerified {
		return ctx.Status(409).JSON(fiber.Map{
			"success": false,
			"message": "User already verified",
		})
	}

	data, err := models.GetUser(user.ID, ctx.Context())
	if err != nil {
		log.Println("Error getting user: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error getting user data",
		})
	}

	token, err := cookies.GenerateEmailToken(user)
	if err != nil {
		log.Println("Error Generating email token: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error generating token",
		})
	}

	err = cookies.SendEmail(token, models.UserPayload{
		Username: user.Username,
		Email:    data.Email,
	})
	if err != nil {
		log.Println("Error sending email: ", err.Error())
		return ctx.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Error sending email to " + data.Email,
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Verification Link Sent to " + data.Email,
	})
}
