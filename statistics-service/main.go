package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "statistics-service/internal/infrastructure/db"
    "statistics-service/internal/infrastructure/nats"
    grpcServer "statistics-service/internal/adapter/grpc"
    natsConsumer "statistics-service/internal/adapter/nats"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Не удалось загрузить .env:", err)
    }

    uri := os.Getenv("MONGO_URI")
    dbName := os.Getenv("DB_NAME")

    database := db.ConnectMongo(uri, dbName)
    nc := nats.ConnectNATS("nats://localhost:4222")
    defer nc.Close()

    go natsConsumer.Start(nc, database)

    if err := grpcServer.StartGRPCServer(database); err != nil {
        log.Fatalf("Failed to start gRPC server: %v", err)
    }
}
