package handler

import (
	"net/http"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	listProductUseCase *usecase.ListProductUseCase
	getProductUseCase  *usecase.GetProductUseCase
}

func NewProductHandler(listProductUseCase *usecase.ListProductUseCase, getProductUseCase *usecase.GetProductUseCase) *ProductHandler {
	return &ProductHandler{
		listProductUseCase: listProductUseCase,
		getProductUseCase:  getProductUseCase,
	}
}

// ListProducts godoc
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "data: array of products"
// @Failure 500 {object} map[string]string "error: error message"
// @Router /products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) {
	products, err := h.listProductUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get product details by product ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} map[string]interface{} "data: product object"
// @Failure 400 {object} map[string]string "error: validation error"
// @Failure 500 {object} map[string]string "error: error message"
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must not be empty",
		})
		return
	}

	result, err := h.getProductUseCase.Execute(usecase.ProductInputDTO{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}
