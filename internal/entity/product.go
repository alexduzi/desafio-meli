package entity

import (
	"fmt"
	"math/rand"
	"time"
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
	Thumbnail   string    `json:"thumbnail" db:"thumbnail"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewProduct(title, description string, price float64, currency, condition string, stock int, sellerID, sellerName, category string) (*Product, error) {
	newId := generateProductID()
	now := time.Now()

	product := &Product{
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
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.Title == "" {
		return fmt.Errorf("title is required")
	}

	if p.Price < 0 {
		return fmt.Errorf("price must be greater than or equal to 0")
	}

	if p.Currency == "" {
		return fmt.Errorf("currency is required")
	}

	if p.Condition != New && p.Condition != Used && p.Condition != Refurbished {
		return fmt.Errorf("condition must be 'new', 'used', or 'refurbished'")
	}

	if p.Stock < 0 {
		return fmt.Errorf("stock must be greater than or equal to 0")
	}

	if p.SellerID == "" {
		return fmt.Errorf("seller_id is required")
	}

	return nil
}

func generateProductID() string {
	timestamp := time.Now().UnixNano()
	random := rand.Int63n(999999)
	return fmt.Sprintf("PROD-%d-%06d", timestamp, random)
}

type ProductImage struct {
	ID           int    `json:"id" db:"id"`
	ProductID    string `json:"product_id" db:"product_id"`
	ImageURL     string `json:"image_url" db:"image_url"`
	DisplayOrder int    `json:"display_order" db:"display_order"`
}

func NewProductImage(productID, imageURL string, displayOrder int) (*ProductImage, error) {
	image := &ProductImage{
		ProductID:    productID,
		ImageURL:     imageURL,
		DisplayOrder: displayOrder,
	}

	if err := image.Validate(); err != nil {
		return nil, err
	}

	return image, nil
}

func (p *ProductImage) Validate() error {
	if p.ProductID == "" {
		return fmt.Errorf("product_id is required")
	}

	if p.ImageURL == "" {
		return fmt.Errorf("image_url is required")
	}

	if p.DisplayOrder < 0 {
		return fmt.Errorf("display_order must be greater than or equal to 0")
	}

	return nil
}
