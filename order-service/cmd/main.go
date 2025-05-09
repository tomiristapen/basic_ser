package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"order-service/config"
	grpcserver "order-service/grpc"
	resthandler "order-service/internal/handler"
	natsinfra "order-service/infrastructure/nats"
	"order-service/internal/repository"
	"order-service/internal/usecase"
	orderpb "order-service/proto"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	inventorypb "inventory-service/proto"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	config.ConnectDB()
	db := config.DB

	repo := repository.NewOrderMongoRepo(db)
	natsConn := natsinfra.ConnectNATS("nats://localhost:4222")
	defer natsConn.Close()

	// Подключение к inventory-service
	inventoryConn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться к inventory-service: %v", err)
	}
	defer inventoryConn.Close()

	inventoryClient := inventorypb.NewInventoryServiceClient(inventoryConn)

	// Создание UseCase с inventoryClient
	uc := usecase.NewOrderUseCase(repo, inventoryClient, natsConn)

	// gRPC сервер
	go func() {
		lis, err := net.Listen("tcp", ":50052")
		if err != nil {
			log.Fatalf("Ошибка запуска gRPC: %v", err)
		}
		s := grpc.NewServer()
		orderpb.RegisterOrderServiceServer(s, grpcserver.NewOrderServer(uc, inventoryClient))
		reflection.Register(s)
		log.Println("gRPC-сервер запущен на :50052")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Ошибка запуска gRPC сервера: %v", err)
		}
	}()

	// REST сервер
	r := gin.Default()
	handler := resthandler.NewOrderHandler(uc)
	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders/:id", handler.GetOrderByID)
	r.PATCH("/orders/:id/status", handler.UpdateOrderStatus)
	r.GET("/orders", handler.GetOrdersByUser)

	port := os.Getenv("ORDER_REST_PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("REST-сервер запущен на :" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Ошибка запуска REST сервера:", err)
	}
}
