package repository

import (
	"context"
	"project/internal/entity"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryInterface interface {
	ListProducts(ctx context.Context) ([]entity.Product, error)
	GetProduct(ctx context.Context, id string) (*entity.Product, error)
	FindImagesByProductID(ctx context.Context, productID string) ([]entity.ProductImage, error)
}

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) ListProducts(ctx context.Context) ([]entity.Product, error) {
	args := m.Called(ctx)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Product), nil
}

func (m *MockProductRepository) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	args := m.Called(ctx, id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), nil
}

func (m *MockProductRepository) FindImagesByProductID(ctx context.Context, productID string) ([]entity.ProductImage, error) {
	args := m.Called(ctx, productID)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.ProductImage), nil
}
