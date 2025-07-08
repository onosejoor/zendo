package oauth_config

import (
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
	clientURL := c.Get("Origin")

	token, err := conf.Exchange(c.Context(), code)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	profile, err := ConvertToken(token.AccessToken)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	isReturned, _ := auth_controllers.HandleOauth(c, *profile)
	if !isReturned {
		return c.Redirect(clientURL)
	}

	return c.Status(500).JSON(fiber.Map{
		"success": false, "message": "Internal server error",
	})
}
