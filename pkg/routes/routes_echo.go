package routes

import (
	"net/http"

	"github.com/amrahman90/go-CRUD-api-sample/pkg/config"

	echo "github.com/labstack/echo/v4"
)

type SetupOptEcho struct {
	Config *config.Config
}

func SetupEcho(opt SetupOptEcho) (app *echo.Echo) {
	app = echo.New()

	{

		itemGroup := app.Group("item")
		{
			// item := new(controllers.UserController)
			itemGroup.GET("/", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.POST("/", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{
					"msg": "ok all",
				})
			})
			itemGroup.GET("/name/:name", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.GET("/:id", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{
					"msg": "ok1",
				})
			})
			itemGroup.PUT("/:id", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{
					"msg": "ok",
				})
			})
			itemGroup.DELETE("/", func(c echo.Context) error {
				return c.JSON(http.StatusOK, map[string]string{
					"msg": "ok",
				})
			})

		}
	}
	return

}

func sendInternalServerErrorEcho(c echo.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Meta: Meta{
			Code:    "ERROR",
			Message: msg,
		},
	})
}

func sendSuccessEcho(c echo.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Meta: Meta{
			Code:    "SUCCESS",
			Message: msg,
		},
		Data: data,
	})
}
