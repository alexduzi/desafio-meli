package usecase

import (
	"fmt"
	"time"
)

type ProductImageDTO struct {
	ID           int    `json:"id" example:"1"`
	ProductID    string `json:"product_id" example:"MLB001"`
	ImageURL     string `json:"image_url" example:"https://images.unsplash.com/photo-1696446702230-a8ff49103cd1?w=800"`
	DisplayOrder int    `json:"display_order" example:"0"`
}

type ProductDTO struct {
	ID          string            `json:"id" example:"MLB001"`
	Title       string            `json:"title" example:"iPhone 15 Pro Max 256GB - Titanium Blue"`
	Description string            `json:"description" example:"Latest Apple flagship smartphone with A17 Pro chip"`
	Price       float64           `json:"price" example:"1299.99"`
	Currency    string            `json:"currency" example:"USD"`
	Condition   string            `json:"condition" example:"new"`
	Stock       int               `json:"stock" example:"45"`
	SellerID    string            `json:"seller_id" example:"SELLER001"`
	SellerName  string            `json:"seller_name" example:"TechWorld Store"`
	Category    string            `json:"category" example:"Electronics > Smartphones"`
	Images      []ProductImageDTO `json:"images,omitempty"`
	Thumbnail   string            `json:"thumbnail,omitempty" example:"https://images.unsplash.com/photo-1696446702230-a8ff49103cd1?w=800"`
	CreatedAt   time.Time         `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time         `json:"updated_at" example:"2024-01-01T00:00:00Z"`
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
		return nil, fmt.Errorf("failed to list products: %w", err)
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
			Thumbnail:   product.Thumbnail,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	return result, nil
}
