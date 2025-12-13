package handler

import (
	"context"
	"net/http"
	"project/internal/dto"

	"github.com/gin-gonic/gin"
)

type ListProductUseCase interface {
	Execute(ctx context.Context) ([]dto.ProductDTO, error)
}

type GetProductUseCase interface {
	Execute(ctx context.Context, input dto.ProductInputDTO) (*dto.ProductDTO, error)
}

type ProductHandler struct {
	listProductUseCase ListProductUseCase
	getProductUseCase  GetProductUseCase
}

func NewProductHandler(listProductUseCase ListProductUseCase, getProductUseCase GetProductUseCase) *ProductHandler {
	return &ProductHandler{
		listProductUseCase: listProductUseCase,
		getProductUseCase:  getProductUseCase,
	}
}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products with thumbnails
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} dto.ProductListResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /api/v1/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	result, err := h.listProductUseCase.Execute(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get product details by product ID including all images
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" example(MLB001)
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	result, err := h.getProductUseCase.Execute(c.Request.Context(), dto.ProductInputDTO{ID: id})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
