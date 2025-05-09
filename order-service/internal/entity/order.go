package entity

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Order struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    string             `bson:"user_id" json:"user_id"`
    Products  []OrderItem        `bson:"products" json:"products"`
    Status    string             `bson:"status" json:"status"`
    CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

type OrderItem struct {
    ProductID string  `bson:"productId" json:"productId"`
    Quantity  int     `bson:"quantity" json:"quantity"`
    Price     float64 `bson:"price" json:"price"` 
}

