package usecase

import (
	"context"

	"inventory-service/internal/entity"
	"inventory-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product *entity.Product) error
	GetAllProducts(ctx context.Context, filter map[string]interface{}, limit, offset int64) ([]entity.Product, error)
	GetProductByID(ctx context.Context, id primitive.ObjectID) (*entity.Product, error)
	UpdateProduct(ctx context.Context, id primitive.ObjectID, product *entity.Product) error
	DeleteProduct(ctx context.Context, id primitive.ObjectID) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repo: r,
	}
}

func (uc *productUseCase) CreateProduct(ctx context.Context, product *entity.Product) error {
	return uc.repo.Create(ctx, product)
}

func (uc *productUseCase) GetAllProducts(ctx context.Context, filter map[string]interface{}, limit, offset int64) ([]entity.Product, error) {
	return uc.repo.GetAll(ctx, filter, limit, offset)
}

func (uc *productUseCase) GetProductByID(ctx context.Context, id primitive.ObjectID) (*entity.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *productUseCase) UpdateProduct(ctx context.Context, id primitive.ObjectID, product *entity.Product) error {
	return uc.repo.Update(ctx, id, product)
}

func (uc *productUseCase) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
	return uc.repo.Delete(ctx, id)
}
