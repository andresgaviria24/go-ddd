package service

import (
	"log"
	"net/http"
	"ws_restaurant/domain/dto"
	"ws_restaurant/infrastructure/persistence"
	"ws_restaurant/infrastructure/repository"
)

type RestaurantServiceImpl struct {
	foodRepository repository.FoodRepository
}

func InitRestaurantServiceImpl() *RestaurantServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &RestaurantServiceImpl{
		foodRepository: dbHelper.FoodRepository,
	}
}

func (r *RestaurantServiceImpl) GetFoods() ([]dto.FoodDto, dto.Response) {

	var response dto.Response
	//var foodsDto []dto.FoodDto

	foods, err := r.foodRepository.GetFoods()

	if err != nil {
		response.Status = http.StatusBadRequest
	}

	foodsDto := dto.FoodDto{}.TransformListEntityToDto(foods)
	/*if ok, resp := utils.ValidateQueryError(err,
		"NO_ROLES_FOUND", "ERROR_GETTING_ROLES", headers.Language); !ok {
		return userRoles, resp
	}

	result, err := r.relationUserRepository.FindUsersByRolAndTenantAndServiceDesk(
		rol.IdRol,
		headers.TenantId,
		idServiceDesk,
	)

	if ok, resp := utils.ValidateQueryError(err,
		"NO_USERS_FOUND", "ERROR_GETTING_USERS", headers.Language); !ok {
		return userRoles, resp
	}

	if len(result) <= 0 {
		response.Status = http.StatusNotFound
		response.Message = utils.Language(headers.Language, "NO_USERS_FOUND")
		return userRoles, response
	}

	userRoles.NameRol = rol.RolName

	for _, user := range result {
		userRoles.Users = append(userRoles.Users, dto.UserDto{
			IdUser:       user.IdUser,
			FullNameUser: user.UserName,
			Email:        user.Email,
		})
	}*/

	response.Status = http.StatusOK
	return foodsDto, response
}
