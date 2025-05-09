package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"user-service/internal/handler"
	db "user-service/internal/infrastructure/db"
	natsinfra "user-service/internal/infrastructure/nats"
	userpb "user-service/proto"
)

func main() {
	// Загрузка переменных окружения
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env:", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	// Подключение к Mongo
	database := db.ConnectMongo(mongoURI, dbName)

	// Подключение к NATS
	nc := natsinfra.ConnectNATS("nats://localhost:4222")
	defer nc.Close()

	// Запуск gRPC сервера
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка запуска TCP: %v", err)
	}

	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, handler.NewUserHandler(database, nc))
	reflection.Register(s)

	log.Println("gRPC-сервер юзер-сервиса запущен на :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Ошибка запуска gRPC сервера: %v", err)
	}
}
