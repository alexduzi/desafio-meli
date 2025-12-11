package usecase

import (
	"fmt"
	"time"
)

type ProductImageDTO struct {
	ID           int    `json:"id"`
	ProductID    string `json:"product_id"`
	ImageURL     string `json:"image_url"`
	DisplayOrder int    `json:"display_order"`
}

type ProductDTO struct {
	ID          string            `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Price       float64           `json:"price"`
	Currency    string            `json:"currency"`
	Condition   string            `json:"condition"`
	Stock       int               `json:"stock"`
	SellerID    string            `json:"seller_id"`
	SellerName  string            `json:"seller_name"`
	Category    string            `json:"category"`
	Images      []ProductImageDTO `json:"images,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type ListProductUseCase struct {
	ProductRepository ProductRepositoryInterface
}

func NewListProductUseCase(productRepo ProductRepositoryInterface) *ListProductUseCase {
	return &ListProductUseCase{
		ProductRepository: productRepo,
	}
}

func (p *ListProductUseCase) Execute() ([]ProductDTO, error) {
	products, err := p.ProductRepository.ListProducts()
	if err != nil {
		return nil, fmt.Errorf("error listing products use case %w", err)
	}

	result := make([]ProductDTO, 0, len(products))

	for _, product := range products {
		result = append(result, ProductDTO{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			Currency:    product.Currency,
			Condition:   product.Condition,
			Stock:       product.Stock,
			SellerID:    product.SellerID,
			SellerName:  product.SellerName,
			Category:    product.Category,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	return result, nil
}
