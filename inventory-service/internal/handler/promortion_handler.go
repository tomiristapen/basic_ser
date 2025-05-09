package handler

import (
    "inventory-service/internal/entity"
    "inventory-service/internal/usecase"
    "net/http"

    "github.com/gin-gonic/gin"
)

type PromotionHandler struct {
    uc usecase.PromotionUseCase
}

func NewPromotionHandler(uc usecase.PromotionUseCase) *PromotionHandler {
    return &PromotionHandler{uc: uc}
}

func (h *PromotionHandler) CreatePromotion(c *gin.Context) {
    var promotion entity.Promotion
    if err := c.ShouldBindJSON(&promotion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.uc.CreatePromotion(c.Request.Context(), promotion); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create promotion"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "promotion created"})
}

func (h *PromotionHandler) GetAllPromotions(c *gin.Context) {
    promotions, err := h.uc.GetAllPromotions(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch promotions"})
        return
    }

    c.JSON(http.StatusOK, promotions)
}

func (h *PromotionHandler) DeletePromotion(c *gin.Context) {
    id := c.Param("id")

    if err := h.uc.DeletePromotion(c.Request.Context(), id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot delete promotion"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "promotion deleted"})
}