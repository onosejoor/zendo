package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	logData := fmt.Sprintf("Request made: %s %s", string(c.Request().Header.Method()), string(c.Request().URI().PathOriginal()))
	fmt.Println(logData)

	return c.Next()
}
