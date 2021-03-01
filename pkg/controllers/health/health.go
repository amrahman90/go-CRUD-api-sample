package health

import (
	"net/http"

	fiber "github.com/gofiber/fiber/v2"
)

type HealthController struct{}

func (h HealthController) Status(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(map[string]string{"status": "working"})
}
