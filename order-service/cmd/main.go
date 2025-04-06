package main

import (
    "order-service/config"
    "order-service/internal/handler"
    "order-service/internal/repository"
    "order-service/internal/usecase"

    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()
    db := config.DB

    repo := repository.NewOrderMongoRepo(db)
    uc := usecase.NewOrderUseCase(repo)
    h := handler.NewOrderHandler(uc)

    r := gin.Default()

    r.POST("/orders", h.CreateOrder)
    r.GET("/orders/:id", h.GetOrderByID)
    r.PATCH("/orders/:id", h.UpdateOrderStatus)
    r.GET("/orders", h.GetOrdersByUser)

    r.Run(":8081") 
}
