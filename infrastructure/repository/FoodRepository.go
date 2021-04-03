/************************
*
* Author: Andrés Gaviria
* Created Date: 10/03/2021
*
* - Se crea metodo de busqueda de usuarios con rol admin de mesa
* - User search method with table admin role is created
 */

package repository

import (
	"ws_restaurant/domain/entity"
)

type FoodRepository interface {
	GetFoods() ([]entity.Food, error)
}
