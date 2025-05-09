package entity

type UserStats struct {
    UserID       string  `bson:"user_id"`
    RegisteredAt string  `bson:"registered_at"`
    OrderCount   int32   `bson:"order_count"`
    TotalSpent   float64 `bson:"total_spent"`
}
