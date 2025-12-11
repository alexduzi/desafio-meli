package database

import (
	"database/sql"
	"fmt"
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
	rows, err := p.DB.Query("SELECT * FROM products")
	if err != nil {
		return nil, fmt.Errorf("error listing products: %w", err)
	}
	defer rows.Close()

	products := make([]entity.Product, 0)

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.ID,
			&product.Title,
			&product.Description,
			&product.Price,
			&product.Currency,
			&product.Condition,
			&product.Stock,
			&product.SellerID,
			&product.SellerName,
			&product.Category,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

func (p *ProductRepository) GetProduct(id string) (*entity.Product, error) {
	row := p.DB.QueryRow("SELECT * FROM products WHERE id = ?", id)
	if row.Err() != nil {
		return nil, fmt.Errorf("error get product %w", row.Err())
	}
	var product entity.Product

	row.Scan(&product.ID, &product.Title, &product.Description,
		&product.Price, &product.Currency, &product.Condition, &product.Stock,
		&product.SellerID, &product.SellerName, &product.Category,
		&product.CreatedAt, &product.UpdatedAt)

	return &product, nil
}

func (p *ProductRepository) FindImagesByProductID(productID string) ([]entity.ProductImage, error) {
	rows, err := p.DB.Query("SELECT * FROM product_images WHERE product_id = ?", productID)
	if err != nil {
		return nil, fmt.Errorf("error find images %w", err)
	}
	defer rows.Close()

	images := []entity.ProductImage{}

	for rows.Next() {
		var pImage entity.ProductImage
		err := rows.Scan(&pImage.ID, &pImage.ProductID, &pImage.ImageURL, &pImage.DisplayOrder)
		if err != nil {
			return nil, fmt.Errorf("error scanning product image: %w", err)
		}
		images = append(images, pImage)
	}

	return images, nil
}
