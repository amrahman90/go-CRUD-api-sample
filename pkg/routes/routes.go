package routes

import (
	"net/http"

	"github.com/amrahman90/go-CRUD-api-sample/pkg/config"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/controllers/health"

	// "github.com/gin-gonic/gin"
	fiber "github.com/gofiber/fiber/v2"
)

type SetupOpt struct {
	Config *config.Config
}

func Setup(opt SetupOpt) (app *fiber.App) {
	// logger := opt.Config.GetLogger()

	// gin.SetMode(gin.ReleaseMode)
	app = fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	// app.Use(ginzap.RecoveryWithZap(logger, true))
	// app.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	health := new(health.HealthController)

	app.Get("/health", health.Status)

	{
		itemGroup := app.Group("item")
		{
			// item := new(controllers.UserController)
			itemGroup.Get("/", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.Post("/", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.Get("/name/:name", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.Get("/id/:id", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.Put("/:id", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.Delete("/", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok",
				})
			})

		}
	}
	return

}
