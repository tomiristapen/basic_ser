package handler

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	userpb "user-service/proto"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	DB   *mongo.Database
	NATS *nats.Conn
}

func NewUserHandler(db *mongo.Database, nc *nats.Conn) *UserHandler {
	return &UserHandler{
		DB:   db,
		NATS: nc,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	userID := uuid.New().String()
	createdAt := time.Now().Format(time.RFC3339)

	_, err := h.DB.Collection("users").InsertOne(ctx, bson.M{
		"user_id":    userID,
		"created_at": createdAt,
	})
	if err != nil {
		return nil, err
	}

	// Событие в NATS
	event := map[string]string{
		"user_id":    userID,
		"created_at": createdAt,
	}
	data, _ := json.Marshal(event)
	h.NATS.Publish("user.created", data)

	return &userpb.CreateUserResponse{
		UserId:    userID,
		CreatedAt: createdAt,
	}, nil
}
