package transactiondto

import "backend/models"

type TransactionResponse struct {
	ID      int                         `json:"id" gorm:"primary_key:auto_increment"`
	Product models.ProductResponse      `json:"product" gorm:"foreignKey:ProductID"`
	Buyer   models.UsersProfileResponse `json:"buyer"`
	Seller  models.UsersProfileResponse `json:"seller"`
	Price   int                         `json:"price"`
	Status  string                      `json:"status" gorm:"type: varchar(25)"`
}
