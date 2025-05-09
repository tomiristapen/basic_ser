package main

import (
	"inventory-service/config"
	grpcHandler "inventory-service/internal/grpc" 
	"inventory-service/internal/handler"
	"inventory-service/internal/repository"
	"inventory-service/internal/routes"
	"inventory-service/internal/usecase"
	pb "inventory-service/proto"

	"log"
	"net"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.ConnectDB()
	db := config.DB

	productCollection := db.Collection("products")
	productRepo := repository.NewProductMongoRepo(productCollection)
	productUC := usecase.NewProductUseCase(productRepo)
	productHandler := handler.NewProductHandler(productUC)
	grpcProductServer := grpcHandler.NewInventoryServer(productUC)

	promotionCollection := db.Collection("promotions")
	promotionRepo := repository.NewPromotionMongoRepo(promotionCollection)
	promotionUC := usecase.NewPromotionUseCase(promotionRepo)
	promotionHandler := handler.NewPromotionHandler(promotionUC)

	go func() {
		r := gin.Default()
		routes.SetupRoutes(r, productHandler, promotionHandler)
		log.Println("REST-сервер запущен на :8080")
		if err := r.Run(":8080"); err != nil {
			log.Fatalf("Ошибка запуска REST-сервера: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50054") // Убедитесь, что порт не занят
	if err != nil {
		log.Fatalf("Ошибка открытия порта для gRPC: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, grpcProductServer)

	reflection.Register(grpcServer)

	log.Println("gRPC-сервер запущен на :50054")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска gRPC-сервера: %v", err)
	}
}
