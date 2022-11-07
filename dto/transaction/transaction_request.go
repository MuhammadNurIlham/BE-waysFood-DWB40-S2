package transactiondto

type TransactionRequest struct {
	ProductID int `gorm:"type: int" json:"productId" validate:"required"`
	SellerID  int `gorm:"type: int" json:"sellerId" validate:"required"`
	Price     int `gorm:"type:int" json:"price" validate:"required"`
}
