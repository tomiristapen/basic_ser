// package entity

// import "time"

// type Promotion struct {
// 	ID                 string `bson:"_id,omitempty" json:"id"`
// 	Name               string
// 	Description        string
// 	DiscountPercentage float64
// 	ApplicableProducts []string `bson:"applicable_products" json:"applicable_products"`
// 	StartDate          time.Time
// 	EndDate            time.Time
// 	IsActive           bool
// }
package entity

import "time"

type Promotion struct {
	ID                 string    `bson:"_id,omitempty" json:"id"`
	Name               string    `bson:"name" json:"name"`
	Description        string    `bson:"description" json:"description"`
	DiscountPercentage float64   `bson:"discountPercentage" json:"discountPercentage"`
	ApplicableProducts []string  `bson:"applicableProducts" json:"applicableProducts"`
	StartDate          time.Time `bson:"startDate" json:"startDate"`
	EndDate            time.Time `bson:"endDate" json:"endDate"`
	IsActive           bool      `bson:"isActive" json:"isActive"`
}
