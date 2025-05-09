package repository

import (
    "context"
    "inventory-service/internal/entity"
)

type PromotionRepository interface {
    Create(ctx context.Context, promotion entity.Promotion) error
    GetAll(ctx context.Context) ([]entity.Promotion, error)
    Delete(ctx context.Context, id string) error
}