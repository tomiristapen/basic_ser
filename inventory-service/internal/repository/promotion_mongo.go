package repository

import (
    "context"
    "inventory-service/internal/entity"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type promotionMongoRepo struct {
    collection *mongo.Collection
}

func NewPromotionMongoRepo(col *mongo.Collection) PromotionRepository {
    return &promotionMongoRepo{collection: col}
}

func (r *promotionMongoRepo) Create(ctx context.Context, promotion entity.Promotion) error {
    _, err := r.collection.InsertOne(ctx, promotion)
    return err
}

func (r *promotionMongoRepo) GetAll(ctx context.Context) ([]entity.Promotion, error) {
    var promotions []entity.Promotion
    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var promotion entity.Promotion
        if err := cursor.Decode(&promotion); err != nil {
            return nil, err
        }
        promotions = append(promotions, promotion)
    }
    return promotions, nil
}

func (r *promotionMongoRepo) Delete(ctx context.Context, id string) error {
    _, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
    return err
}