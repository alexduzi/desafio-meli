package usecase

import (
	"fmt"
)

type ProductDTO struct {
	ID string `json:"id"`
}

type ListProductUseCase struct {
	ProductRepository ProductRepositoryInterface
}

func NewListProductUseCase(productRepo ProductRepositoryInterface) *ListProductUseCase {
	return &ListProductUseCase{
		ProductRepository: productRepo,
	}
}

func (p *ListProductUseCase) Execute() ([]ProductDTO, error) {
	products, err := p.ProductRepository.ListProducts()
	if err != nil {
		return nil, fmt.Errorf("error when listing products %w", err)
	}

	result := make([]ProductDTO, 0, len(products))

	for _, product := range products {
		result = append(result, ProductDTO{
			ID: product.ID,
		})
	}

	return result, nil
}
