package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Admin_id     string             `json:"admin_id,"omitempty" bson:"admin_id,omitempty"`
	First_name   string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	Lat_name     string             `json:"last_name,"omitempty"bson:"last_name,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	Role         string             `json:"role,omitempty" bson:"role,omitempty"`
	created_time time.Time          `json:"created_time,omitempty" bson:"created_time,omitempty"`
	updated_time time.Time          `json:"updated_time,omitempty" bson:"updated_time,omitempty"`
}
