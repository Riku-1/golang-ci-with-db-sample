package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{db: db}
}

func (r ProductRepository) Get(id uint) (Product, error) {
	var product Product
	err := r.db.First(&product, id).Error

	return product, err
}

func (r ProductRepository) Create(code string, price uint) (Product, error) {
	product := Product{Code: code, Price: price}
	err := r.db.Create(&product).Error

	return product, err
}

func (r ProductRepository) Update(id uint, product Product) (Product, error) {
	_product, err := r.Get(id)
	if err != nil {
		return Product{}, err
	}

	err = r.db.Model(&_product).Updates(product).Error

	return product, err
}

func (r ProductRepository) Delete(id uint) error {
	var product Product
	err := r.db.Delete(&product, id).Error

	return err
}
