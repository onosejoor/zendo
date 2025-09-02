package routes

import (
	oauth_config "main/configs/oauth"
	"main/handlers/auth_controllers"
	"main/handlers/email_controllers"
	"main/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	oauth := oauth_config.InitializeOauthConfig()

	auth := app.Group("/auth")

	auth.Get("/user", middlewares.AuthMiddleware, auth_controllers.HandleGetUser)
	auth.Put("/user", middlewares.AuthMiddleware, auth_controllers.UpdateUserController)
	auth.Get("/verify_email", email_controllers.HandleVerifyEmailController)
	auth.Post("/verify_email", middlewares.AuthMiddleware, email_controllers.SendEmailTokenController)
	auth.Get("/refresh-token", auth_controllers.HandleAccessToken)
	auth.Post("/signup", auth_controllers.HandleSignup)
	auth.Post("/signin", auth_controllers.HandleSignin)

	auth.Get("/oauth/google", oauth.GetOauthController)
	auth.Get("/oauth/callback", oauth.OauthCallBackController)
	auth.Post("/oauth/exchange", oauth.OauthExchangeController)
}
