package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Staff struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Staff_id     string             `json:"staff_id,omitempty" bson:"staff_id,omitempty"`
	First_name   string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	Last_name    string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Department   string             `json:"department,omitempty" bson:"department,omitempty"`
	Age          int                `json:"age,omitempty" bson:"age,omitempty"`
	Gender       string             `json:"gender,omitempty" bson:"gender,omitempty"`
	Salary       float64            `json:"salary,omitempty" bson:"salary,omitempty"`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	Created_time time.Time          `json:"created_time,omitempty" bson:"created_time,omitempty"`
	Updated_time time.Time          `json:"updated_time,omitempty" bson:"updated_time,omitempty"`
}
