package usecase

import (
	"project/internal/dto"
	"project/internal/entity"
)

func toListProductDTO(products []entity.Product) []dto.ProductDTO {
	result := make([]dto.ProductDTO, 0, len(products))

	for _, product := range products {
		result = append(result, dto.ProductDTO{
			ID:        product.ID,
			Title:     product.Title,
			Price:     product.Price,
			Currency:  product.Currency,
			Condition: product.Condition,
			Stock:     product.Stock,
			Category:  product.Category,
			Thumbnail: product.Thumbnail,
		})
	}

	return result
}

func toGetProductDTO(product entity.Product, images []entity.ProductImage) *dto.ProductDTO {
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
	}
}
