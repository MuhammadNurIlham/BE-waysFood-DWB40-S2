package productdto

import "backend/models"

type ProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Qty   int    `json:"qty"`
	User  string `json:"user"`
}

type ProductTransactionRequest struct {
	ID    int                         `json:"id"`
	Name  string                      `json:"name"`
	Desc  string                      `json:"desc"`
	Price int                         `json:"price"`
	Image string                      `json:"image"`
	Qty   int                         `json:"qty"`
	User  models.UsersProfileResponse `json:"user"`
}
