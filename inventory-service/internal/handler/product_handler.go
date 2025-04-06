package handler

import (
    "context"
    "inventory-service/internal/entity"
    "inventory-service/internal/usecase"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type ProductHandler struct {
    uc usecase.ProductUseCase
}

func NewProductHandler(uc usecase.ProductUseCase) *ProductHandler {
    return &ProductHandler{uc: uc}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var product entity.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err := h.uc.CreateProduct(ctx, &product)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    products, err := h.uc.GetAllProducts(ctx)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, products)
}
