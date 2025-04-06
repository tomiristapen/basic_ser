package repository

import (
    "context"
    "inventory-service/internal/entity"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type productMongo struct {
    collection *mongo.Collection
}

func NewProductMongoRepo(db *mongo.Database) ProductRepository {
    return &productMongo{
        collection: db.Collection("products"),
    }
}

func (r *productMongo) Create(ctx context.Context, product *entity.Product) error {
    _, err := r.collection.InsertOne(ctx, product)
    return err
}

func (r *productMongo) GetAll(ctx context.Context) ([]entity.Product, error) {
    var products []entity.Product

    cursor, err := r.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var p entity.Product
        if err := cursor.Decode(&p); err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    return products, nil
}
