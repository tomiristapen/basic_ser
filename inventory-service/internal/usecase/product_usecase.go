package usecase

import (
    "context"
    "inventory-service/internal/entity"
    "inventory-service/internal/repository"
)


import "go.mongodb.org/mongo-driver/bson/primitive"

func (uc *productUseCase) CreateProduct(ctx context.Context, product *entity.Product) error {
    product.ID = primitive.NewObjectID() 
    return uc.repo.Create(ctx, product)
}

type ProductUseCase interface {
    CreateProduct(ctx context.Context, product *entity.Product) error
    GetAllProducts(ctx context.Context) ([]entity.Product, error)
}

type productUseCase struct {
    repo repository.ProductRepository
}

func NewProductUseCase(r repository.ProductRepository) ProductUseCase {
    return &productUseCase{
        repo: r,
    }
}

// func (uc *productUseCase) CreateProduct(ctx context.Context, product *entity.Product) error {
//     return uc.repo.Create(ctx, product)
// }

func (uc *productUseCase) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
    return uc.repo.GetAll(ctx)
}
