package repository

import (
    "context"
    "order-service/internal/entity"
)

type OrderRepository interface {
    Create(ctx context.Context, order *entity.Order) error
    GetByID(ctx context.Context, id string) (*entity.Order, error)
    UpdateStatus(ctx context.Context, id string, status string) error
    GetByUser(ctx context.Context, userID string) ([]entity.Order, error)
}
