package nats

import (
    "context"
    "encoding/json"
    "log"

    "github.com/nats-io/nats.go"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type UserCreatedEvent struct {
    UserID    string `json:"user_id"`
    CreatedAt string `json:"created_at"`
}

type OrderCreatedEvent struct {
    UserID    string  `json:"user_id"`
    OrderID   string  `json:"order_id"`
    CreatedAt string  `json:"created_at"`
    Amount    float64 `json:"amount"`
}

func Start(nc *nats.Conn, db *mongo.Database) {
    nc.Subscribe("user.created", func(msg *nats.Msg) {
        var event UserCreatedEvent
        if err := json.Unmarshal(msg.Data, &event); err != nil {
            log.Println("Invalid user.created event:", err)
            return
        }

        _, err := db.Collection("user_stats").UpdateOne(
            context.TODO(),
            map[string]interface{}{"user_id": event.UserID},
            map[string]interface{}{
                "$setOnInsert": map[string]interface{}{
                    "user_id":       event.UserID,
                    "registered_at": event.CreatedAt,
                    "order_count":   0,
                    "total_spent":   0,
                },
            },
            options.Update().SetUpsert(true),
        )
        if err != nil {
            log.Println("Failed to upsert user stat:", err)
        } else {
            log.Println("User stat ensured for:", event.UserID)
        }
    })

    nc.Subscribe("order.created", func(msg *nats.Msg) {
        var event OrderCreatedEvent
        if err := json.Unmarshal(msg.Data, &event); err != nil {
            log.Println("Invalid order.created event:", err)
            return
        }

        _, err := db.Collection("user_stats").UpdateOne(
            context.TODO(),
            map[string]interface{}{"user_id": event.UserID},
            map[string]interface{}{
                "$inc": map[string]interface{}{
                    "order_count": 1,
                    "total_spent": event.Amount,
                },
                "$setOnInsert": map[string]interface{}{
                    "registered_at": event.CreatedAt,
                },
            },
            options.Update().SetUpsert(true),
        )

        if err != nil {
            log.Println("Failed to update stats:", err)
        } else {
            log.Println("Updated stats for:", event.UserID)
        }
    })

    log.Println("NATS consumers started")
}
