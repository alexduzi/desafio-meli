package database

import (
	"project/internal/entity"
)

type ProductRepositoryInterface interface {
	ListProducts() ([]entity.Product, error)
}
