package dto

import "ws_restaurant/domain/entity"

type FoodDto struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"prices"`
}

func (g FoodDto) TransformListEntityToDto(f []entity.Food) []FoodDto {
	var result []FoodDto
	for _, fd := range f {
		result = append(result, FoodDto{
			Id:    fd.Id,
			Name:  fd.Name,
			Price: fd.Price,
		})
	}
	return result
}
