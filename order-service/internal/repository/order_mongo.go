package repository

import (
    "context"
    "order-service/internal/entity"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type orderMongo struct {
    collection *mongo.Collection
}

func NewOrderMongoRepo(db *mongo.Database) OrderRepository {
    return &orderMongo{
        collection: db.Collection("orders"),
    }
}

func (r *orderMongo) Create(ctx context.Context, order *entity.Order) error {
    _, err := r.collection.InsertOne(ctx, order)
    return err
}

func (r *orderMongo) GetByID(ctx context.Context, id string) (*entity.Order, error) {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    var order entity.Order
    err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&order)
    if err != nil {
        return nil, err
    }
    return &order, nil
}

func (r *orderMongo) UpdateStatus(ctx context.Context, id string, status string) error {
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    update := bson.M{"$set": bson.M{"status": status}}
    _, err = r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
    return err
}

func (r *orderMongo) GetByUser(ctx context.Context, userID string) ([]entity.Order, error) {
    var orders []entity.Order

    cursor, err := r.collection.Find(ctx, bson.M{"userId": userID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var order entity.Order
        if err := cursor.Decode(&order); err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }

    return orders, nil
}
