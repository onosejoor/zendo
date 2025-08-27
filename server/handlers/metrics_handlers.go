package handlers

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(c *fiber.Ctx) error {
	expectedUsername := os.Getenv("METRICS_USERNAME")
	expectedPassword := os.Getenv("METRICS_PASSWORD")

	if expectedUsername != "" && expectedPassword != "" {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			c.Set("WWW-Authenticate", `Basic realm="metrics", charset="UTF-8"`)
			return c.Status(401).SendString("Unauthorized")
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")
		decodedBytes, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			c.Set("WWW-Authenticate", `Basic realm="metrics", charset="UTF-8"`)
			return c.Status(401).SendString("Unauthorized")
		}

		credentials := string(decodedBytes)
		parts := strings.SplitN(credentials, ":", 2)
		if len(parts) != 2 {
			c.Set("WWW-Authenticate", `Basic realm="metrics", charset="UTF-8"`)
			return c.Status(401).SendString("Unauthorized")
		}

		username := parts[0]
		password := parts[1]

		if username != expectedUsername || password != expectedPassword {
			c.Set("WWW-Authenticate", `Basic realm="metrics", charset="UTF-8"`)
			return c.Status(401).SendString("Unauthorized")
		}
	}

	handler := adaptor.HTTPHandler(promhttp.Handler())
	return handler(c)
}
