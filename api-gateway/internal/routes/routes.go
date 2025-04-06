package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// продукты
	router.GET("/products", handler.HandleProductProxy)
	router.GET("/products/:id", handler.HandleProductProxy)

	router.POST("/products", middleware.AuthRequired(), handler.HandleProductProxy)
	router.PATCH("/products/:id", middleware.AuthRequired(), handler.HandleProductProxy)
	router.DELETE("/products/:id", middleware.AuthRequired(), handler.HandleProductProxy)

	//заказы
	order := router.Group("/orders", middleware.AuthRequired())
	{
		order.POST("", handler.HandleOrderProxy)
		order.PATCH("/:id", handler.HandleOrderProxy)
		order.GET("", handler.HandleOrderProxy)
		order.GET("/:id", handler.HandleOrderProxy)
	}
}
