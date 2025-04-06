package repository

import (
	"context"
	"inventory-service/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	GetAll(ctx context.Context, filter map[string]interface{}, limit, offset int64) ([]entity.Product, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Product, error)
	Update(ctx context.Context, id primitive.ObjectID, product *entity.Product) error
	Delete(ctx context.Context, id primitive.ObjectID) error
}
