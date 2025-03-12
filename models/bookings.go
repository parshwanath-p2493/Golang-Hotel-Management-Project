package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BookingId    string             `json:"bookingid,omitempty" bson:"bookingid,omitempty"`
	Guest_id     string             `json:"guest_id" bson:"guest_id"`
	Room_id      string             `json:"room_id" bson:"room_id"`
	Room_number  int32              `json:"room_number" bson:"room_number"`
	Status       string             `json:"status" bson:"status"`
	Food_Items   []string           `json:"food_items" bson:"food_items"`
	BookingDate  time.Time          `json:"bookingdate" bson:"bookingdate"`
	CheckOutDate time.Time          `json:"checkoutdate" bson:"checkoutdate"`
	Created_time time.Time          `json:"created_time" bson:"created_time"`
	Updated_time time.Time          `json:"updated_time" bson:"updated_time"`
}
