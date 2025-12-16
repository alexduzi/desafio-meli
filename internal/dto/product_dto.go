package dto

import (
	"time"
)

type ProductInputDTO struct {
	ID string `json:"id"`
}

type ProductImageDTO struct {
	ID           int    `json:"id" example:"1"`
	ProductID    string `json:"product_id" example:"MLB001"`
	ImageURL     string `json:"image_url" example:"https://images.unsplash.com/photo-1696446702230-a8ff49103cd1?w=800"`
	DisplayOrder int    `json:"display_order" example:"0"`
}

type ProductDTO struct {
	ID          string            `json:"id" example:"MLB001"`
	Title       string            `json:"title" example:"iPhone 15 Pro Max 256GB - Titanium Blue"`
	Description string            `json:"description,omitempty" example:"Latest Apple flagship smartphone with A17 Pro chip"`
	Price       float64           `json:"price" example:"1299.99"`
	Currency    string            `json:"currency" example:"USD"`
	Condition   string            `json:"condition" example:"new"`
	Stock       int               `json:"stock" example:"45"`
	SellerID    string            `json:"seller_id,omitempty" example:"SELLER001"`
	SellerName  string            `json:"seller_name,omitempty" example:"TechWorld Store"`
	Category    string            `json:"category" example:"Electronics > Smartphones"`
	Images      []ProductImageDTO `json:"images,omitempty"`
	Thumbnail   string            `json:"thumbnail,omitempty" example:"https://images.unsplash.com/photo-1696446702230-a8ff49103cd1?w=800"`
	CreatedAt   time.Time         `json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt   time.Time         `json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
}

type ProductListResponse struct {
	Data []ProductDTO `json:"data"`
}

type ProductResponse struct {
	Data ProductDTO `json:"data"`
}
