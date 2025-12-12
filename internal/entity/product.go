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
	ID          string    `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Price       float64   `json:"price" db:"price"`
	Currency    string    `json:"currency" db:"currency"`
	Condition   string    `json:"condition" db:"condition"`
	Stock       int       `json:"stock" db:"stock"`
	SellerID    string    `json:"seller_id" db:"seller_id"`
	SellerName  string    `json:"seller_name" db:"seller_name"`
	Category    string    `json:"category" db:"category"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
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
	ID           int    `json:"id" db:"id"`
	ProductID    string `json:"product_id" db:"product_id"`
	ImageURL     string `json:"image_url" db:"image_url"`
	DisplayOrder int    `json:"display_order" db:"display_order"`
}
