package oauth_config

import (
	"fmt"
	"log"
	"main/handlers/auth_controllers"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthConfig struct {
	*oauth2.Config
}

func InitializeOauthConfig() *OauthConfig {
	return &OauthConfig{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("G_CLIENT_ID"),
			ClientSecret: os.Getenv("G_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("G_REDIRECT"),
			Endpoint:     google.Endpoint,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		},
	}
}

func (conf *OauthConfig) GetOauthController(c *fiber.Ctx) error {

	URL := conf.AuthCodeURL("")

	return c.Redirect(URL)
}

func (conf *OauthConfig) OauthCallBackController(c *fiber.Ctx) error {
	code := c.Query("code")
	clientURL := os.Getenv("FRONTEND_URL")

	url := fmt.Sprintf("%v/auth/oauth-callback?code=%s", clientURL, code)

	return c.Redirect(url)

}

func (conf *OauthConfig) OauthExchangeController(c *fiber.Ctx) error {
	code := c.Query("code")

	token, err := conf.Exchange(c.Context(), code)
	if err != nil {
		log.Println("Error Exchanging Ouath Token: ", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false, "message": err.Error(),
		})
	}

	profile, err := ConvertToken(token.AccessToken)
	if err != nil {
		log.Println("Error Converting Ouath Token: ", err)
		return c.Status(500).JSON(fiber.Map{
			"success": false, "message": "Error Exchanging token",
		})
	}

	isReturned, _ := auth_controllers.HandleOauth(c, *profile)
	if !isReturned {
		return c.Status(200).JSON(fiber.Map{
			"success": true, "message": "Oauth Successful, welcome " + profile.Email,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"success": false, "message": "Internal server error",
	})
}
