package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Managers struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	First_name    string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	Last_name     string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	Manager_id    string             `json:"manager_id,omitempty" bson:"manager_id,omitempty"`
	Department    string             `json:"department,omitempty" bson:"department,omitempty"`
	Age           int32              `json:"age,omitempty" bson:"age,omitempty"`
	Phone         int32              `json:"phone,omitempty" bson:"phone,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	Password      string             `json:"password,omitempty" bson:"password,omitempty"`
	Token         string             `bson:"token,omitempty" json:"token,omitempty"`
	Refresh_Token string             `bson:"refresh_token,omitempty" json:"refresh_token,omitempty"`
	Salary        string             `json:"salary,omitempty" bson:"salary,omitempty"`
	Created_time  time.Time          `json:"created_time,omitempty" bson:"created_time,omitempty"`
	Updated_time  time.Time          `json:"updated_time,omitempty" bson:"updated_time,omitempty"`
}
	