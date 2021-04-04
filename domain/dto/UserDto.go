/*
*
* Auhtor: Héctor Castellanos
* Created Date: 25/02/2021
*
* - struct para representar la información del usuario
* - struct that represents the user information
 */

package dto

type UserDto struct {
	IdUser       string `json:"idUser"`
	FullNameUser string `json:"fullNameUser"`
	Email        string `json:"email"`
}
