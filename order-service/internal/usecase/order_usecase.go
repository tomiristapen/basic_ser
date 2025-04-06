package usecase

import (
    "context"
    "order-service/internal/entity"
    "order-service/internal/repository"
    "time"
)


import "go.mongodb.org/mongo-driver/bson/primitive"

func (uc *orderUseCase) CreateOrder(ctx context.Context, order *entity.Order) error {
    order.ID = primitive.NewObjectID() 
    order.Status = "pending"
    order.CreatedAt = time.Now()
    return uc.repo.Create(ctx, order)
}

type OrderUseCase interface {
    CreateOrder(ctx context.Context, order *entity.Order) error
    GetOrderByID(ctx context.Context, id string) (*entity.Order, error)
    UpdateOrderStatus(ctx context.Context, id string, status string) error
    GetOrdersByUser(ctx context.Context, userID string) ([]entity.Order, error)
}

type orderUseCase struct {
    repo repository.OrderRepository
}

func NewOrderUseCase(r repository.OrderRepository) OrderUseCase {
    return &orderUseCase{
        repo: r,
    }
}

func (uc *orderUseCase) GetOrderByID(ctx context.Context, id string) (*entity.Order, error) {
    return uc.repo.GetByID(ctx, id)
}

func (uc *orderUseCase) UpdateOrderStatus(ctx context.Context, id string, status string) error {
    return uc.repo.UpdateStatus(ctx, id, status)
}

func (uc *orderUseCase) GetOrdersByUser(ctx context.Context, userID string) ([]entity.Order, error) {
    return uc.repo.GetByUser(ctx, userID)
}
