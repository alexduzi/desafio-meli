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
