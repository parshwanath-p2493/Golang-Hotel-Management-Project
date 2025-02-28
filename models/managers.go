package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Managers struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	First_name   string
	Last_name    string
	Manager_id   string
	Department   string
	Age          int32
	Phone        int32
	Email        string
	Password     string
	salary       string
	created_time time.Time
	updated_time time.Time
}
