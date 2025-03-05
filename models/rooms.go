package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rooms struct {
	ID                  primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Room_id             string             `json:"room_id,omitempty" bson:"room_id,omitempty"`
	//Manager_id          string             `json:"manager_id,omitempty" bson:"manager_id,omitempty"`
	Room_number         int32              `json:"room_number,omitempty" bson:"room_number,omitempty"`
	Room_type           string             `json:"room_type,omitempty" bson:"room_type,omitempty"`
	Capacity            string             `json:"capacity,omitempty" bson:"capacity,omitempty"`
	Availability_status string             `json:"availability_status,omitempty" bson:"availability_status,omitempty"`
	Price               float64            `json:"price,omitempty" bson:"price,omitempty"`
	Created_time        time.Time          `json:"created_time,omitempty" bson:"created_time,omitempty"`
	Updated_time        time.Time          `json:"updated_time,omitempty" bson:"updated_time,omitempty"`
}
