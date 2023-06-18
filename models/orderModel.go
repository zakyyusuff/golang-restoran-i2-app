package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	Order_Date time.Time          `json:"order_date" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Order_id   string             `json:"order_id"`
	Table_id   *string            `json:"table_id" validate:"required"`
}

// ///////////////////////////////////////////////// coba
// package models

// import (
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type Order struct {
// 	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
// 	OrderDate time.Time          `bson:"order_date" json:"order_date" validate:"required"`
// 	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
// 	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
// 	OrderID   string             `bson:"order_id" json:"order_id"`
// 	TableID   *string            `bson:"table_id" json:"table_id" validate:"required"`
// }
