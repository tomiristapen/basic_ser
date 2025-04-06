package handler

import (
    "context"
    "net/http"
    "order-service/internal/entity"
    "order-service/internal/usecase"
    "time"

    "github.com/gin-gonic/gin"
)

type OrderHandler struct {
    uc usecase.OrderUseCase
}

func NewOrderHandler(uc usecase.OrderUseCase) *OrderHandler {
    return &OrderHandler{uc: uc}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
    var order entity.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := h.uc.CreateOrder(ctx, &order); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
    id := c.Param("id")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    order, err := h.uc.GetOrderByID(ctx, id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
    id := c.Param("id")
    var input struct {
        Status string `json:"status"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := h.uc.UpdateOrderStatus(ctx, id, input.Status); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "status updated"})
}

func (h *OrderHandler) GetOrdersByUser(c *gin.Context) {
    userID := c.Query("user")
    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "missing user ID"})
        return
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    orders, err := h.uc.GetOrdersByUser(ctx, userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, orders)
}
