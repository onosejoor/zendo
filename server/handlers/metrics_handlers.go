package handlers

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MetricsHandler(c *fiber.Ctx) error {
	handler := adaptor.HTTPHandler(promhttp.Handler())
	return handler(c)
}
