package fooditem

import (
	"github.com/amrahman90/go-CRUD-api-sample/pkg/db"
	"gorm.io/gorm"
)

type FoodItemService struct {
	DB *gorm.DB
}

func NewFoodItem(opt FoodItemService) *FoodItemService {
	return &opt
}

func (s FoodItemService) GetAllItems() (items []db.FoodItem, err error) {
	return
}

func (s FoodItemService) GetItem() (items db.FoodItem, err error) {
	return
}

func (s FoodItemService) AddItem() (err error) {
	return
}

func (s FoodItemService) UpdateItem() (err error) {
	return
}

func (s FoodItemService) DeleteItem() (err error) {
	return
}
