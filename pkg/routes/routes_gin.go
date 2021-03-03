package routes

import (
	"net/http"
	"time"

	"github.com/amrahman90/go-CRUD-api-sample/pkg/config"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/services/fooditem"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/utils"
	"go.uber.org/zap"

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
		DB:     dbConn,
		Logger: logger,
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
			itemGroup.GET("/", func(c *gin.Context) {
				opt := fooditem.GetAllItemsOpt{
					Limit:  utils.StrToInt(c.DefaultQuery("limit", "10")),
					Offset: utils.StrToInt(c.DefaultQuery("offset", "0")),
				}

				result, err := fooditemService.GetAllItems(&opt)
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", result)
			})
			itemGroup.POST("/", func(c *gin.Context) {
				opt := fooditem.FoodItem{}
				err := c.ShouldBindJSON(&opt)
				if err != nil {
					logger.Error("AddItem", zap.Error(err))
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				err = fooditemService.AddItem(opt)
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", nil)
			})
			itemGroup.GET("/name/:name", func(c *gin.Context) {
				result, err := fooditemService.GetItem(fooditem.GetItemOpt{
					Name: c.Param("name"),
				})
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", result)
			})
			itemGroup.GET("/id/:id", func(c *gin.Context) {
				result, err := fooditemService.GetItem(fooditem.GetItemOpt{
					ID: utils.StrToInt(c.Param("id")),
				})
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", result)
			})
			itemGroup.PUT("/:id", func(c *gin.Context) {
				item := fooditem.FoodItem{}
				err := c.ShouldBindJSON(&item)
				if err != nil {
					logger.Error("AddItem", zap.Error(err))
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				err = fooditemService.UpdateItem(fooditem.UpdateItemOpt{
					ID:   utils.StrToInt(c.Param("id")),
					Item: item,
				})
				if err != nil {
					sendInternalServerErrorGin(c, "some error occurred")
					return
				}
				sendSuccessGin(c, "Success", nil)
			})
			itemGroup.DELETE("/:id", func(c *gin.Context) {
				err := fooditemService.DeleteItem(utils.StrToInt(c.Param("id")))
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
