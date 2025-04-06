package routes

import (
	"inventory-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *handler.ProductHandler) {
	r.POST("/products", h.CreateProduct)
	r.GET("/products", h.GetAllProducts)
	r.GET("/products/:id", h.GetProductByID)
	r.PATCH("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)
}
