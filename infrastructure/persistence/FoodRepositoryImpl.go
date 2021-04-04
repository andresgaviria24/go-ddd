package persistence

import (
	"ws_restaurant/domain/entity"

	"github.com/jinzhu/gorm"
)

type FoodRepositoryImpl struct {
	db *gorm.DB
}

func (r *FoodRepositoryImpl) GetFoods() ([]entity.Food, error) {

	var listUserDetail []entity.Food
	err := r.db.Find(&listUserDetail).Error

	return listUserDetail, err
}
