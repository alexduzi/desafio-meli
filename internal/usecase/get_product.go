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
	ProductRepository repository.ProductRepositoryInterface
}

func NewGetProductUseCase(productRepo repository.ProductRepositoryInterface) *GetProductUseCase {
	return &GetProductUseCase{
		ProductRepository: productRepo,
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

	product, err := p.ProductRepository.GetProduct(ctx, input.ID)
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

	images, err := p.ProductRepository.FindImagesByProductID(ctx, input.ID)
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

	imagesDto := make([]dto.ProductImageDTO, 0, len(images))
	for _, image := range images {
		imagesDto = append(imagesDto, dto.ProductImageDTO{
			ID:           image.ID,
			ProductID:    image.ProductID,
			ImageURL:     image.ImageURL,
			DisplayOrder: image.DisplayOrder,
		})
	}

	return &dto.ProductDTO{
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
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
		Images:      imagesDto,
	}, nil
}
