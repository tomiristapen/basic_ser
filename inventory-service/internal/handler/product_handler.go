package handler

import (
	"inventory-service/internal/entity"
	"inventory-service/internal/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandler struct {
	uc usecase.ProductUseCase
}

func NewProductHandler(u usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{uc: u}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.CreateProduct(c.Request.Context(), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	filter := make(map[string]interface{})
	if category := c.Query("category"); category != "" {
		filter["category"] = category
	}
	if minPrice := c.Query("minPrice"); minPrice != "" {
		if val, err := strconv.ParseFloat(minPrice, 64); err == nil {
			filter["minPrice"] = val
		}
	}
	if maxPrice := c.Query("maxPrice"); maxPrice != "" {
		if val, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			filter["maxPrice"] = val
		}
	}

	
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	offset, _ := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 64)

	products, err := h.uc.GetAllProducts(c.Request.Context(), filter, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	product, err := h.uc.GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.uc.UpdateProduct(c.Request.Context(), id, &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated"})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.uc.DeleteProduct(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
