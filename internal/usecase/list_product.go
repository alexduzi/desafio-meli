package usecase

import (
	"fmt"
	"project/internal/dto"
	"project/internal/repository"
)

type ListProductUseCase struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewListProductUseCase(productRepo repository.ProductRepositoryInterface) *ListProductUseCase {
	return &ListProductUseCase{
		ProductRepository: productRepo,
	}
}

func (p *ListProductUseCase) Execute() ([]dto.ProductDTO, error) {
	products, err := p.ProductRepository.ListProducts()
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}

	result := make([]dto.ProductDTO, 0, len(products))

	for _, product := range products {
		result = append(result, dto.ProductDTO{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			Currency:    product.Currency,
			Condition:   product.Condition,
			Stock:       product.Stock,
			SellerID:    product.SellerID,
			SellerName:  product.SellerName,
			Category:    product.Category,
			Thumbnail:   product.Thumbnail,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	return result, nil
}
