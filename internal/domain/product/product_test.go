package product

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	productName        = "Product 1"
	productDescription = "Product Description"
)

func Test_NewProduct_Create(t *testing.T) {
	assert := assert.New(t)

	product, err := NewProduct(productName, productDescription)

	assert.Nil(err)
	assert.NotNil(product)
	assert.Equal(productName, product.Name)
	assert.Equal(productDescription, product.Description)
}

func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	product, err := NewProduct(productName, productDescription)

	assert.Nil(err)
	assert.NotNil(product)
	assert.Equal(productName, product.Name)
	assert.Greater(product.CreatedAt, now)
}
