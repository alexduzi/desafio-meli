package usecase

import (
	"project/internal/entity"
)

type ProductRepositoryInterface interface {
	ListProducts() ([]entity.Product, error)
	GetProduct(id string) (*entity.Product, error)
}
