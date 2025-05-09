package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
    ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name     string             `bson:"name" json:"name"`
    Category string             `bson:"category" json:"category"`
    Stock    int                `bson:"stock" json:"stock"`
    Price    float64            `bson:"price" json:"price"`
}
