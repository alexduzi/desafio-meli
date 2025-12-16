package usecase

import (
	"context"
	"fmt"
	"project/internal/dto"
	"project/internal/errors"
	"project/internal/repository"
	"strings"

	"github.com/rs/zerolog/log"
)

type GetProductUseCase struct {
	productRepository repository.ProductRepositoryInterface
}

func NewGetProductUseCase(productRepo repository.ProductRepositoryInterface) *GetProductUseCase {
	return &GetProductUseCase{
		productRepository: productRepo,
	}
}

func (p *GetProductUseCase) Execute(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error) {
	log.Debug().
		Str("product_id", input.ID).
		Msg("Executing GetProduct use case")

	if strings.TrimSpace(input.ID) == "" {
		log.Warn().Msg("Invalid product ID: empty or whitespace")
		return nil, errors.ErrInvalidProductID
	}

	product, err := p.productRepository.GetProduct(ctx, input.ID)
	if err != nil {
		log.Error().
			Err(err).
			Str("product_id", input.ID).
			Msg("Failed to get product from repository")
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	log.Debug().
		Str("product_id", input.ID).
		Str("product_title", product.Title).
		Msg("Product found successfully")

	images, err := p.productRepository.FindImagesByProductID(ctx, input.ID)
	if err != nil {
		log.Error().
			Err(err).
			Str("product_id", input.ID).
			Msg("Failed to get product images")
		return nil, fmt.Errorf("failed to get product images: %w", err)
	}

	log.Debug().
		Str("product_id", input.ID).
		Int("images_count", len(images)).
		Msg("Product images retrieved")

	return toGetProductDTO(*product, images), nil
}
