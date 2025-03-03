package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID            primitive.ObjectID
	Food_id       string
	Item_name     string
	Category_name string
	Category_id   string
	Price         float64 `json:"price","omitempty" bson:"price","omitempty"`
	created_time  time.Time
	updated_time  time.Time
}
