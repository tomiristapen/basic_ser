package routes

import (
	"inventory-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, productHandler *handler.ProductHandler, promotionHandler *handler.PromotionHandler) {
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.GetAllProducts)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.PATCH("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)
 
	r.POST("/promotions", promotionHandler.CreatePromotion)
	r.GET("/promotions", promotionHandler.GetAllPromotions)
	r.DELETE("/promotions/:id", promotionHandler.DeletePromotion)
}
