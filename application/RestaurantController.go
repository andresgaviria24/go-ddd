package application

import (
	"net/http"
	"ws_restaurant/domain/dto"
	"ws_restaurant/domain/service"

	"github.com/gin-gonic/gin"
)

type RestaurantController struct {
	restaurantService service.RestaurantService
}

func InitRestaurantController(router *gin.Engine) {
	restaurantController := RestaurantController{
		restaurantService: service.InitRestaurantServiceImpl(),
	}
	router.GET("/food", restaurantController.GetFoodcHandler)
}

func (r *RestaurantController) GetFoodcHandler(c *gin.Context) {

	var response dto.Response

	foods, response := r.restaurantService.GetFoods()

	if response.Status != http.StatusOK {
		c.JSON(response.Status, response)
		return
	}
	c.JSON(response.Status, foods)
}
