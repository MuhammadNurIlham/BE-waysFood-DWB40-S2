package productdto

type ProductRequest struct {
	Name   string `json:"name" form:"name" gorm:"type: varchar(255)" validate:"required"`
	Desc   string `json:"desc" gorm:"type:text" form:"desc" validate:"required"`
	Price  int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Image  string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty    int    `json:"qty" form:"qty" gorm:"type: int" validate:"required"`
	UserID int    `json:"user_id" form:"user_id"`
}
