package database

import (
	"database/sql"
	"fmt"
	"project/internal/entity"
	"project/internal/errors"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (p *ProductRepository) ListProducts() ([]entity.Product, error) {
	products := []entity.Product{}

	query := "SELECT * FROM products ORDER BY created_at DESC"

	err := p.DB.Select(&products, query)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDatabaseError, err)
	}

	return products, nil
}

func (p *ProductRepository) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product

	query := "SELECT * FROM products WHERE id = ?"

	err := p.DB.Get(&product, query, id)
	if err == sql.ErrNoRows {
		return nil, errors.ErrProductNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrDatabaseError, err)
	}

	return &product, nil
}

func (p *ProductRepository) FindImagesByProductID(productID string) ([]entity.ProductImage, error) {
	images := []entity.ProductImage{}

	query := "SELECT * FROM product_images WHERE product_id = ? ORDER BY display_order ASC"

	err := p.DB.Select(&images, query, productID)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("%w: %v", errors.ErrDatabaseError, err)
	}

	return images, nil
}
