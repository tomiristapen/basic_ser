package usecase

import (
    "context"
    "inventory-service/internal/entity"
    "inventory-service/internal/repository"
)

type PromotionUseCase interface {
    CreatePromotion(ctx context.Context, promotion entity.Promotion) error
    GetAllPromotions(ctx context.Context) ([]entity.Promotion, error)
    DeletePromotion(ctx context.Context, id string) error
}

type promotionUseCase struct {
    repo repository.PromotionRepository
}

func NewPromotionUseCase(r repository.PromotionRepository) PromotionUseCase {
    return &promotionUseCase{repo: r}
}

func (uc *promotionUseCase) CreatePromotion(ctx context.Context, promotion entity.Promotion) error {
    return uc.repo.Create(ctx, promotion)
}

func (uc *promotionUseCase) GetAllPromotions(ctx context.Context) ([]entity.Promotion, error) {
    return uc.repo.GetAll(ctx)
}

func (uc *promotionUseCase) DeletePromotion(ctx context.Context, id string) error {
    return uc.repo.Delete(ctx, id)
}