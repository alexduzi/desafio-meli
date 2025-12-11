package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	productTitle       = "Product 1"
	productDescription = "A great product"
	price              = 100.50
	currency           = "BRL"
	stock              = 10
	sellerID           = "seller123"
	sellerName         = "John Doe"
	category           = "Electronics"
)

func Test_NewProduct_Create(t *testing.T) {
	assert := assert.New(t)

	product, err := NewProduct(productTitle, productDescription, price, currency, New, stock, sellerID, sellerName, category)

	assert.Nil(err)
	assert.NotNil(product)
	assert.Equal(productTitle, product.Title)
	assert.Equal(productDescription, product.Description)
	assert.Equal(price, product.Price)
	assert.Equal(currency, product.Currency)
	assert.Equal(New, product.Condition)
	assert.Equal(stock, product.Stock)
	assert.Equal(sellerID, product.SellerID)
	assert.Equal(sellerName, product.SellerName)
	assert.Equal(category, product.Category)
	assert.NotEmpty(product.ID)
	assert.False(product.CreatedAt.IsZero())
	assert.False(product.UpdatedAt.IsZero())
}
