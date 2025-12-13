package usecase

import (
	"context"
	"fmt"
	"project/internal/dto"
	"project/internal/errors"
	"project/internal/repository"
	"strings"
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
	if strings.TrimSpace(input.ID) == "" {
		return nil, errors.ErrInvalidProductID
	}

	product, err := p.ProductRepository.GetProduct(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	images, err := p.ProductRepository.FindImagesByProductID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product images: %w", err)
	}

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
