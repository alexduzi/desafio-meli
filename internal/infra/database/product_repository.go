package database

import (
	"database/sql"
	"project/internal/entity"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) ListProducts() ([]entity.Product, error) {
	products := make([]entity.Product, 0, 5)

	product, err := entity.NewProduct("Product test", entity.New, entity.Price{Amount: 90.57, Currency: "BRL"}, 10, 3)
	if err != nil {
		return nil, err
	}

	products = append(products, *product)

	return products, nil
}

func (p *ProductRepository) GetProduct(id string) (*entity.Product, error) {
	return &entity.Product{
		ID: id,
	}, nil
}
