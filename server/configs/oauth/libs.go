package oauth_config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/configs"

	"net/http"
)

// this is what the data would look like, after the convert
// you can use this reference to get the data and save to db for example.

func ConvertToken(accessToken string) (*configs.GooglePayload, error) {
	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if httpErr != nil {
		return nil, httpErr
	}

	defer resp.Body.Close()

	// Reads the entire HTTP body from resp.Body using ioutil.ReadAll.
	// If any error occurs during the read operation, it is
	// returned as bodyErr. Otherwise it is stored in the respBody variable.
	respBody, bodyErr := io.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	var body map[string]any
	if err := json.Unmarshal(respBody, &body); err != nil {
		return nil, err
	}

	// If json body containing error,
	// then the token is indeed invalid. return invalid token err
	if body["error"] != nil {
		return nil, errors.New("invalid token")
	}

	var data configs.GooglePayload
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	if !data.EmailVerified {
		return nil, errors.New("email not verified")

	}

	return &data, nil
}
