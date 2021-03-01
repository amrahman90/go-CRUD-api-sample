package routes

import (
	"net/http"
	"time"

	"github.com/amrahman90/go-CRUD-api-sample/pkg/config"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/services/fooditem"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

type SetupOptGin struct {
	Config *config.Config
}

func SetupGin(opt SetupOptGin) (app *gin.Engine, err error) {
	logger := opt.Config.GetLogger()
	dbConn, err := opt.Config.GetDB()
	if err != nil {
		return
	}
	fooditemService := fooditem.NewFoodItem(fooditem.FoodItemService{
		DB: dbConn,
	})

	gin.SetMode(gin.ReleaseMode)
	app = gin.New()

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	app.Use(ginzap.RecoveryWithZap(logger, true))
	app.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	{
		itemGroup := app.Group("item")
		{
			// item := new(controllers.UserController)
			itemGroup.GET("/", func(c *gin.Context) {
				result, err := fooditemService.GetAllItems()
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", result)
			})
			itemGroup.POST("/", func(c *gin.Context) {
				err := fooditemService.AddItem()
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", nil)
			})
			itemGroup.GET("/name/:name", func(c *gin.Context) {
				result, err := fooditemService.GetItem()
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", result)
			})
			itemGroup.GET("/id/:id", func(c *gin.Context) {
				result, err := fooditemService.GetItem()
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", result)
			})
			itemGroup.PUT("/:id", func(c *gin.Context) {
				err := fooditemService.UpdateItem()
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", nil)
			})
			itemGroup.DELETE("/", func(c *gin.Context) {
				err := fooditemService.UpdateItem()
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", nil)
			})

		}
	}
	return

}

func sendInternalServerErrorGin(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Meta: Meta{
			Code:    "ERROR",
			Message: msg,
		},
	})
}

func sendSuccessGin(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Meta: Meta{
			Code:    "SUCCESS",
			Message: msg,
		},
		Data: data,
	})
}

type Meta struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}
