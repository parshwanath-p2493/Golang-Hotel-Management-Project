package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Guest struct {
	ID            primitive.ObjectID `bson:"_id"`
	Guest_id      string             `bson:"guest_id,omitempty" json:"guest_id,omitempty"`
	ID_Proof_Type string             `bson:"id_proof_type,omitempty" json:"id_proof_type,omitempty"`
	First_Name    string             `bson:"first_name,omitempty" json:"first_name,omitempty"`
	Last_Name     string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Phone         int                `bson:"phone,omitempty" json:"phone,omitempty"`
	Email         string             `bson:"email,omitempty" json:"email,omitempty"`
	Password      string             `bson:"Password,omitempty" json:"Password,omitempty"`
	Gender        string             `bson:"gender,omitempty" json:"gender,omitempty"`
	Country       string
	created_time  time.Time
	updated_time  time.Time
}
