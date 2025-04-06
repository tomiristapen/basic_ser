package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/products", handler.HandleProductProxy)
	router.POST("/products", handler.HandleProductProxy) 

	order := router.Group("/orders")
	{
		order.POST("", middleware.AuthRequired(), handler.HandleOrderProxy)
		order.PATCH("/:id", middleware.AuthRequired(), handler.HandleOrderProxy)

		// Открытые маршруты
		order.GET("", handler.HandleOrderProxy)       // ?user=...
		order.GET("/:id", handler.HandleOrderProxy)   // по ID
	}
}
