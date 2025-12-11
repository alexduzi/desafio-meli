package entity

import (
	"time"

	"github.com/rs/xid"
)

const (
	New         = "new"
	Used        = "used"
	Refurbished = "refurbished"
)

type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	Condition   string    `json:"condition"`
	Stock       int       `json:"stock"`
	SellerID    string    `json:"seller_id"`
	SellerName  string    `json:"seller_name"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProduct(title, description string, price float64, currency, condition string, stock int, sellerID, sellerName, category string) (*Product, error) {
	newId := xid.New().String()
	now := time.Now()

	return &Product{
		ID:          newId,
		Title:       title,
		Description: description,
		Price:       price,
		Currency:    currency,
		Condition:   condition,
		Stock:       stock,
		SellerID:    sellerID,
		SellerName:  sellerName,
		Category:    category,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

type ProductImage struct {
	ID           int    `json:"id"`
	ProductID    string `json:"product_id"`
	ImageURL     string `json:"image_url"`
	DisplayOrder int    `json:"display_order"`
}
