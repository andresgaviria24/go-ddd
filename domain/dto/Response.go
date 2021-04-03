package dto

type Response struct {
	Status      int    `json:"status"`
	Description string `json:"description"`
	Message     string `json:"message"`
}
