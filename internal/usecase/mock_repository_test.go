package usecase

import (
	"project/internal/entity"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) ListProducts() ([]entity.Product, error) {
	args := m.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Product), nil
}

func (m *MockProductRepository) GetProduct(id string) (*entity.Product, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), nil
}

func (m *MockProductRepository) FindImagesByProductID(productID string) ([]entity.ProductImage, error) {
	args := m.Called(productID)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.ProductImage), nil
}
