package handler

import (
	"net/http"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	listProductUseCase *usecase.ListProductUseCase
}

func NewProductHandler(listProductUseCase *usecase.ListProductUseCase) *ProductHandler {
	return &ProductHandler{
		listProductUseCase: listProductUseCase,
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
