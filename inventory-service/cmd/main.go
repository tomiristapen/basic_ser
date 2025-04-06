package main

import (
	"inventory-service/config"
	"inventory-service/internal/handler"
	"inventory-service/internal/repository"
	"inventory-service/internal/routes"
	"inventory-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	db := config.DB

	productCollection := db.Collection("products")
	productRepo := repository.NewProductMongoRepo(productCollection)
	productUC := usecase.NewProductUseCase(productRepo)
	productHandler := handler.NewProductHandler(productUC)

	r := gin.Default()

	routes.SetupRoutes(r, productHandler)

	r.Run(":8080")
}
