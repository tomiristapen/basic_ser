package main

import (
    "inventory-service/config"
    "inventory-service/internal/handler"
    "inventory-service/internal/repository"
    "inventory-service/internal/usecase"
    "github.com/gin-gonic/gin"
)

func main() {
    // Подключаем БД
    config.ConnectDB()
    db := config.DB

    // Слои
    productRepo := repository.NewProductMongoRepo(db)
    productUC := usecase.NewProductUseCase(productRepo)
    productHandler := handler.NewProductHandler(productUC)

    // Запуск сервера
    r := gin.Default()

    r.POST("/products", productHandler.CreateProduct)
    r.GET("/products", productHandler.GetAllProducts)

    r.Run(":8080") // http://localhost:8080
}
