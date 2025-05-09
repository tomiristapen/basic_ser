package grpc

import (
	"context"
	"log"
	"net"

	"statistics-service/internal/entity"
	pb "statistics-service/proto/statistics"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedStatisticsServiceServer
	DB *mongo.Database
}

func (s *Server) GetUserStatistics(ctx context.Context, req *pb.UserStatisticsRequest) (*pb.UserStatisticsResponse, error) {
	var result entity.UserStats
	err := s.DB.Collection("user_stats").FindOne(ctx, bson.M{"user_id": req.UserId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	// ✅ Логируем регистрацию пользователя
	log.Printf("REGISTERED → user_id: %s | registered_at: %s", result.UserID, result.RegisteredAt)

	return &pb.UserStatisticsResponse{
		UserId:       result.UserID,
		RegisteredAt: result.RegisteredAt,
	}, nil
}

func (s *Server) GetUserOrdersStatistics(ctx context.Context, req *pb.UserOrderStatisticsRequest) (*pb.UserOrderStatisticsResponse, error) {
	var result entity.UserStats
	err := s.DB.Collection("user_stats").FindOne(ctx, bson.M{"user_id": req.UserId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	// ✅ Логируем заказы и сумму
	log.Printf("STATS → user_id: %s | orders: %d | spent: %.2f", result.UserID, result.OrderCount, result.TotalSpent)

	return &pb.UserOrderStatisticsResponse{
		UserId:     result.UserID,
		OrderCount: result.OrderCount,
		TotalSpent: result.TotalSpent,
	}, nil
}

func StartGRPCServer(db *mongo.Database) error {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStatisticsServiceServer(grpcServer, &Server{DB: db})
	reflection.Register(grpcServer)
	log.Println("gRPC server started on port :50053")
	return grpcServer.Serve(lis)
}
