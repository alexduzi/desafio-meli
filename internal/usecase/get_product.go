package usecase

import (
	"fmt"
	"project/internal/errors"
	"strings"
)

type ProductInputDTO struct {
	ID string `json:"id"`
}

type GetProductUseCase struct {
	ProductRepository ProductRepositoryInterface
}

func NewGetProductUseCase(productRepo ProductRepositoryInterface) *GetProductUseCase {
	return &GetProductUseCase{
		ProductRepository: productRepo,
	}
}

func (p *GetProductUseCase) Execute(input ProductInputDTO) (*ProductDTO, error) {
	// Validate input
	if strings.TrimSpace(input.ID) == "" {
		return nil, errors.ErrInvalidProductID
	}

	product, err := p.ProductRepository.GetProduct(input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	images, err := p.ProductRepository.FindImagesByProductID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product images: %w", err)
	}

	imagesDto := make([]ProductImageDTO, 0, len(images))

	for _, image := range images {
		imagesDto = append(imagesDto, ProductImageDTO{
			ID:           image.ID,
			ProductID:    image.ProductID,
			ImageURL:     image.ImageURL,
			DisplayOrder: image.DisplayOrder,
		})
	}

	return &ProductDTO{
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
