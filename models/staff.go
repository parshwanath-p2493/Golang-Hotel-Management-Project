package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Staff struct {
	ID           primitive.ObjectID `json:"id,omitempty"bson:"_id,omitempty"`
	Staff_id     string             `json:"staff_id,omitempty" bson:"staff_id,omitempty"`
	First_name   string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	Last_name    string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Age          int                `json:"age,omitempty"bson:"age,omitempty"`
	Gender       string             `json:"age,omitempty"bson:"age,omitempty"`
	salary       float64            `json:"salary,omitempty bson`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	created_time time.Time          `json:"created_time,omitempty" bson:"created_time,omitempty"`
	updated_time time.Time          `json:"updated_time,omitempty" bson:"updated_time,omitempty"`
}
