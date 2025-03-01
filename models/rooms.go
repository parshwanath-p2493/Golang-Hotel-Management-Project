package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rooms struct {
	ID                  primitive.ObjectID
	Room_id             string
	Manager_id          string
	Room_number         int32
	Room_type           string
	capacity            string
	Availability_status string
	price               float64
	created_time        time.Time
	updated_time        time.Time
}
