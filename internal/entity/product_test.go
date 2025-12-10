package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	productTitle = "Product 1"
	price        = Price{Amount: 100.50, Currency: "BRL"}
	stock        = 10
	soldQuantity = 4
)

func Test_NewProduct_Create(t *testing.T) {
	assert := assert.New(t)

	product, err := NewProduct(productTitle, New, price, stock, soldQuantity)

	assert.Nil(err)
	assert.NotNil(product)
	assert.Equal(productTitle, product.Title)
	assert.Equal(price, product.Price)
	assert.Equal(New, product.Condition)
	assert.Equal(stock, product.Stock)
	assert.Equal(soldQuantity, product.SoldQuantity)
}
