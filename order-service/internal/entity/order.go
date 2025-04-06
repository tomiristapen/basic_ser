package entity

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Order struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    UserID    string             `bson:"userId" json:"userId"`
    Products  []OrderItem        `bson:"products" json:"products"`
    Status    string             `bson:"status" json:"status"` // pending, completed, cancelled
    CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type OrderItem struct {
    ProductID string `bson:"productId" json:"productId"`
    Quantity  int    `bson:"quantity" json:"quantity"`
}
