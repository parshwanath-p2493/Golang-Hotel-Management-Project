package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BookingId    string             `json:"bookingid,omitempty" bson:"bookingid" validate:"required"`
	Guest_id     string             `json:"guest_id" bson:"guest_id" validate:"required"`
	Room_id      string             `json:"room_id" bson:"room_id" validate:"required"`
	Room_number  int32              `json:"room_number" bson:"room_number" validate:"required"`
	Status       string             `json:"status" bson:"status"`
	Food_Items   []string           `json:"food_items" bson:"food_items"`
	Token        string             `bson:"token,omitempty" json:"token,omitempty"`
	BookingDate  time.Time          `json:"bookingdate" bson:"bookingdate" validate:"required"`
	CheckOutDate time.Time          `json:"checkoutdate" bson:"checkoutdate" validate:"required"`
	Created_time time.Time          `json:"created_time" bson:"created_time" validate:"required"`
	Updated_time time.Time          `json:"updated_time" bson:"updated_time" validate:"required"`
}
