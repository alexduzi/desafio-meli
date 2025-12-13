package usecase

import (
	"project/internal/entity"
)

type MockProductRepository struct {
	ListProductsFunc            func() ([]entity.Product, error)
	GetProductFunc              func(id string) (*entity.Product, error)
	FindImagesByProductIDFunc   func(productID string) ([]entity.ProductImage, error)
}

func (m *MockProductRepository) ListProducts() ([]entity.Product, error) {
	if m.ListProductsFunc != nil {
		return m.ListProductsFunc()
	}
	return []entity.Product{}, nil
}

func (m *MockProductRepository) GetProduct(id string) (*entity.Product, error) {
	if m.GetProductFunc != nil {
		return m.GetProductFunc(id)
	}
	return nil, nil
}

func (m *MockProductRepository) FindImagesByProductID(productID string) ([]entity.ProductImage, error) {
	if m.FindImagesByProductIDFunc != nil {
		return m.FindImagesByProductIDFunc(productID)
	}
	return []entity.ProductImage{}, nil
}
