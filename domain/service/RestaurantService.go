package service

import "ws_restaurant/domain/dto"

type RestaurantService interface {
	GetFoods() ([]dto.FoodDto, dto.Response)
}
