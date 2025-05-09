package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"order-service/internal/entity"
	"order-service/internal/repository"
	"time"

	"github.com/nats-io/nats.go"
	inventorypb "inventory-service/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderUseCase interface {
	CreateOrder(ctx context.Context, order *entity.Order) error
	GetOrderByID(ctx context.Context, id string) (*entity.Order, error)
	GetOrdersByUser(ctx context.Context, userID string) ([]entity.Order, error)
	UpdateOrderStatus(ctx context.Context, id, status string) error
}

type orderUseCase struct {
	repo            repository.OrderRepository
	inventoryClient inventorypb.InventoryServiceClient
	nats            *nats.Conn
}

func NewOrderUseCase(r repository.OrderRepository, inv inventorypb.InventoryServiceClient, nc *nats.Conn) OrderUseCase {
	return &orderUseCase{
		repo:            r,
		inventoryClient: inv,
		nats:            nc,
	}
}

func (uc *orderUseCase) CreateOrder(ctx context.Context, order *entity.Order) error {
	order.ID = primitive.NewObjectID()
	order.Status = "pending"
	order.CreatedAt = time.Now()

	// Получаем цену по каждому productId через gRPC
	total := 0.0
	for i, item := range order.Products {
		res, err := uc.inventoryClient.GetProductByID(ctx, &inventorypb.GetProductRequest{
			Id: item.ProductID,
		})
		if err != nil {
			return fmt.Errorf("product not found: %w", err)
		}
		price := res.Price
		order.Products[i].Price = price
		total += price * float64(item.Quantity)
	}

	if err := uc.repo.Create(ctx, order); err != nil {
		return err
	}

	event := map[string]interface{}{
		"user_id":    order.UserID,
		"order_id":   order.ID.Hex(),
		"created_at": order.CreatedAt.Format(time.RFC3339),
		"amount":     total,
	}

	data, _ := json.Marshal(event)
	uc.nats.Publish("order.created", data)

	return nil
}

func (uc *orderUseCase) GetOrderByID(ctx context.Context, id string) (*entity.Order, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *orderUseCase) GetOrdersByUser(ctx context.Context, userID string) ([]entity.Order, error) {
	return uc.repo.GetByUser(ctx, userID)
}

func (uc *orderUseCase) UpdateOrderStatus(ctx context.Context, id, status string) error {
	return uc.repo.UpdateStatus(ctx, id, status)
}
