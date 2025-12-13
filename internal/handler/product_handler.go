package handler

import (
	"net/http"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ListProductUseCase interface {
	Execute() ([]usecase.ProductDTO, error)
}

type GetProductUseCase interface {
	Execute(input usecase.ProductInputDTO) (*usecase.ProductDTO, error)
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

type ProductListResponse struct {
	Data []usecase.ProductDTO `json:"data"`
}

type ProductResponse struct {
	Data usecase.ProductDTO `json:"data"`
}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products with thumbnails
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} ProductListResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /api/v1/products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	result, err := h.listProductUseCase.Execute()
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
// @Success 200 {object} ProductResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 404 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")

	result, err := h.getProductUseCase.Execute(usecase.ProductInputDTO{ID: id})
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
