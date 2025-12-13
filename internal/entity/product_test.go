package entity

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewProduct_Success(t *testing.T) {
	product, err := NewProduct(
		"MLB001",
		"iPhone 15 Pro Max",
		"Latest Apple smartphone with A17 Pro chip",
		1299.99,
		"USD",
		New,
		25,
		"seller-001",
		"Apple Store",
		"Electronics",
	)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "iPhone 15 Pro Max", product.Title)
	assert.Equal(t, 1299.99, product.Price)
	assert.Equal(t, New, product.Condition)
	assert.True(t, strings.HasPrefix(product.ID, "MLB"))
	assert.False(t, product.CreatedAt.IsZero())
	assert.False(t, product.UpdatedAt.IsZero())
}

func Test_NewProduct_AllConditions(t *testing.T) {
	tests := []struct {
		name      string
		condition string
	}{
		{"New product", New},
		{"Used product", Used},
		{"Refurbished product", Refurbished},
	}

	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := NewProduct(fmt.Sprintf("%s%d", "MLB00", idx), "Test", "Desc", 99.99, "USD", tt.condition, 10, "seller-001", "Seller", "Cat")
			assert.NoError(t, err)
			assert.Equal(t, tt.condition, product.Condition)
		})
	}
}

func Test_NewProduct_UniqueIDs(t *testing.T) {
	product1, _ := NewProduct("MLB001", "Product 1", "Desc", 10.0, "USD", New, 5, "seller1", "Seller", "Cat")
	product2, _ := NewProduct("MLB002", "Product 2", "Desc", 20.0, "USD", New, 10, "seller2", "Seller", "Cat")

	assert.NotEqual(t, product1.ID, product2.ID)
}

func Test_NewProduct_ValidationErrors(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		price     float64
		currency  string
		condition string
		stock     int
		sellerID  string
		wantErr   string
	}{
		{"Empty title", "", 99.99, "USD", New, 10, "seller-001", "title is required"},
		{"Negative price", "Product", -10.0, "USD", New, 10, "seller-001", "price must be greater than or equal to 0"},
		{"Empty currency", "Product", 99.99, "", New, 10, "seller-001", "currency is required"},
		{"Invalid condition", "Product", 99.99, "USD", "broken", 10, "seller-001", "condition must be 'new', 'used', or 'refurbished'"},
		{"Negative stock", "Product", 99.99, "USD", New, -5, "seller-001", "stock must be greater than or equal to 0"},
		{"Empty seller ID", "Product", 99.99, "USD", New, 10, "", "seller_id is required"},
	}

	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := NewProduct(fmt.Sprintf("%s%d", "MLB00", idx), tt.title, "Desc", tt.price, tt.currency, tt.condition, tt.stock, tt.sellerID, "Seller", "Cat")
			assert.Error(t, err)
			assert.Nil(t, product)
			assert.Equal(t, tt.wantErr, err.Error())
		})
	}
}

func Test_NewProduct_EdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		description string
		price       float64
		stock       int
	}{
		{"Empty description", "", 99.99, 10},
		{"Zero price", "Free item", 0.0, 10},
		{"Zero stock", "Out of stock", 99.99, 0},
	}

	for idx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := NewProduct(fmt.Sprintf("%s%d", "MLB00", idx), "Product", tt.description, tt.price, "USD", New, tt.stock, "seller-001", "Seller", "Cat")
			assert.NoError(t, err)
			assert.NotNil(t, product)
		})
	}
}

func Test_NewProductImage_Success(t *testing.T) {
	image, err := NewProductImage("PROD-001", "https://cdn.store.com/image.jpg", 0)

	assert.NoError(t, err)
	assert.NotNil(t, image)
	assert.Equal(t, "PROD-001", image.ProductID)
	assert.Equal(t, "https://cdn.store.com/image.jpg", image.ImageURL)
	assert.Equal(t, 0, image.DisplayOrder)
}

func Test_NewProductImage_ValidationErrors(t *testing.T) {
	tests := []struct {
		name         string
		productID    string
		imageURL     string
		displayOrder int
		wantErr      string
	}{
		{"Empty product ID", "", "https://example.com/image.jpg", 0, "product_id is required"},
		{"Empty image URL", "PROD-001", "", 0, "image_url is required"},
		{"Negative display order", "PROD-001", "https://example.com/image.jpg", -1, "display_order must be greater than or equal to 0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			image, err := NewProductImage(tt.productID, tt.imageURL, tt.displayOrder)
			assert.Error(t, err)
			assert.Nil(t, image)
			assert.Equal(t, tt.wantErr, err.Error())
		})
	}
}

func Test_NewProductImage_MultipleOrders(t *testing.T) {
	orders := []int{0, 1, 2, 5, 10}

	for _, order := range orders {
		image, err := NewProductImage("PROD-001", "https://example.com/image.jpg", order)
		assert.NoError(t, err)
		assert.Equal(t, order, image.DisplayOrder)
	}
}

func Test_Product_Constants(t *testing.T) {
	assert.Equal(t, "new", New)
	assert.Equal(t, "used", Used)
	assert.Equal(t, "refurbished", Refurbished)
}
