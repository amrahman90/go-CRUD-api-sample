package routes

import (
	"net/http"

	"github.com/amrahman90/go-CRUD-api-sample/pkg/config"

	// "github.com/gin-gonic/gin"
	fiber "github.com/gofiber/fiber/v2"
)

type SetupOptFiber struct {
	Config *config.Config
}

func SetupFiber(opt SetupOptFiber) (app *fiber.App) {
	// logger := opt.Config.GetLogger()

	// gin.SetMode(gin.ReleaseMode)
	app = fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

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
					"msg": "ok all",
				})
			})
			itemGroup.Get("/name/:name", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.Get("/:id", func(c *fiber.Ctx) error {
				return c.Status(http.StatusOK).JSON(map[string]string{
					"msg": "ok1",
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

func sendInternalServerErrorFiber(c *fiber.Ctx, msg string) error {
	return c.Status(http.StatusInternalServerError).JSON(Response{
		Meta: Meta{
			Code:    "ERROR",
			Message: msg,
		},
	})
}

func sendSuccessFiber(c *fiber.Ctx, msg string, data interface{}) error {
	return c.Status(http.StatusOK).JSON(Response{
		Meta: Meta{
			Code:    "SUCCESS",
			Message: msg,
		},
		Data: data,
	})
}
