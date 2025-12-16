package usecase

import (
	"context"
	"fmt"
	"project/internal/dto"
	"project/internal/repository"

	"github.com/rs/zerolog/log"
)

type ListProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewListProductUseCase(productRepo repository.ProductRepositoryInterface) *ListProductUseCase {
	return &ListProductUseCase{
		productRepository: productRepo,
	}
}

func (p *ListProductUseCase) Execute(ctx context.Context) ([]dto.ProductDTO, error) {
	log.Debug().Msg("Executing ListProducts use case")

	products, err := p.productRepository.ListProducts(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to list products from repository")
		return nil, fmt.Errorf("failed to list products: %w", err)
	}

	log.Info().
		Int("products_count", len(products)).
		Msg("Products listed successfully")

	return toListProductDTO(products), nil
}
