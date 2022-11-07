package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	GetProductUser(UserID int) ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Find(&products).Error

	return products, err
}

func (r *repository) GetProductUser(UserID int) ([]models.Product, error) {
	var product []models.Product
	err := r.db.Where("user_id=?", UserID).Preload("User").Find(&product).Error

	return product, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

// func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
// 	err := r.db.Model(&product).Error
// 	return product, err
// }

func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
