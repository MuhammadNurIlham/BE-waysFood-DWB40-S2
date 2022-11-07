package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `gorm:"type: varchar(255)" form:"gender" json:"gender" validate:"required"`
	Phone    string `gorm:"type: varchar(255)" form:"phone" json:"phone" validate:"required"`
	Role     string `gorm:"type: varchar(255)" form:"role" json:"role" validate:"required"`
	Location string `gorm:"type: varchar(255)" form:"location" json:"location" validate:"required"`
	Image    string `gorm:"type: varchar(255)" form:"image"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Image    string `json:"image" form:"image"`
	Phone    string `json:"phone" form:"phone"`
	Location string `json:"location" form:"location"`
	Role     string `json:"role" form:"role"`
	Gender   string `json:"gender" form:"gender"`
}
