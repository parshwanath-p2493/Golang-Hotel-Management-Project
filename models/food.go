package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Food_id       string             `json:"food_id,omitempty" bson:"food_id,omitempty"`
	Item_name     string             `json:"item_name,omitempty" bson:"item_name,omitempty"`
	Category_name string             `json:"category_name,omitempty" bson:"category_name,omitempty"`
	Category_id   string             `json:"category_id,omitempty" bson:"category_id,omitempty"`
	Price         float64            `json:"price,omitempty" bson:"price,omitempty"`
	Created_time  time.Time          `json:"created_time,omitempty" bson:"created_time,omitempty"`
	Updated_time  time.Time          `json:"updated_time,omitempty" bson:"updated_time,omitempty"`
}
