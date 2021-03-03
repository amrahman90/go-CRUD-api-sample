package fooditem

import (
	"github.com/amrahman90/go-CRUD-api-sample/pkg/db"
	"github.com/amrahman90/go-CRUD-api-sample/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type FoodItemService struct {
	Logger *zap.Logger
	DB     *gorm.DB
}

func NewFoodItem(opt FoodItemService) *FoodItemService {
	return &opt
}

type GetAllItemsOpt struct {
	Limit  int `json:",omitempty"`
	Offset int `json:",omitempty"`
}
type GetAllItemsResp struct {
	Items      []*db.FoodItem
	Pagination struct {
		Limit  int
		Offset int
	}
}

func (s FoodItemService) GetAllItems(opt *GetAllItemsOpt) (resp GetAllItemsResp, err error) {
	fn := "FoodItemService.GetAllItems"
	var items []*db.FoodItem
	if result := s.DB.Offset(opt.Offset).Limit(opt.Limit).Where(db.FoodItem{}).Find(&items); result.Error != nil {
		logger.Error(fn, zap.Error(result.Error))
		err = result.Error
		return
	}
	resp = GetAllItemsResp{
		Items: items,
		Pagination: struct {
			Limit  int
			Offset int
		}{
			Limit:  opt.Limit,
			Offset: opt.Offset + len(items),
		},
	}
	logger.Debug(fn, zap.Any("result", items))
	return
}

type GetItemOpt struct {
	ID   int
	Name string
}

func (s FoodItemService) GetItem(opt GetItemOpt) (item db.FoodItem, err error) {
	fn := "GetItem"
	query := db.FoodItem{}
	if len(opt.Name) > 0 {
		query.Name = opt.Name
	} else {
		query.ID = &opt.ID
	}
	if result := s.DB.Where(query).First(&item); result.Error != nil {
		logger.Error(fn, zap.Error(result.Error))
		err = result.Error
		return
	}
	return
}

type FoodItem struct {
	Name         string
	Description  string
	Price        *float64 `json:"omitempty"`
	RestaurantID *int
}

func (s FoodItemService) AddItem(opt FoodItem) (err error) {
	fn := "AddItem"
	logger := s.Logger

	if result := s.DB.Create(&db.FoodItem{
		Name:         opt.Name,
		Description:  opt.Description,
		Price:        opt.Price,
		RestaurantID: opt.RestaurantID,
	}); result.Error != nil {
		logger.Error(fn, zap.Error(result.Error))
		return result.Error
	}
	return
}

type UpdateItemOpt struct {
	ID   int
	Item FoodItem
}

func (s FoodItemService) UpdateItem(opt UpdateItemOpt) (err error) {
	fn := "UpdateItem"
	logger := s.Logger

	if result := s.DB.Model(&db.FoodItem{
		ID: &opt.ID,
	}).Where(&db.FoodItem{}).Updates(&db.FoodItem{
		Name:         opt.Item.Name,
		Description:  opt.Item.Description,
		Price:        opt.Item.Price,
		RestaurantID: opt.Item.RestaurantID,
	}); result.Error != nil {
		logger.Error(fn, zap.Error(result.Error))
		return result.Error
	}
	return
}

func (s FoodItemService) DeleteItem(id int) (err error) {
	fn := "DeleteItem"
	if result := s.DB.Delete(&db.FoodItem{}, id); result.Error != nil {
		logger.Error(fn, zap.Error(result.Error))
		return result.Error
	}
	return
}
