package repository

import (
    "context"
    "inventory-service/internal/entity"
)

type ProductRepository interface {
    Create(ctx context.Context, product *entity.Product) error
    GetAll(ctx context.Context) ([]entity.Product, error)
}
