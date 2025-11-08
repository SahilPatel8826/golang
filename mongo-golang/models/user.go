package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    string `json:"age" bson:"age"`
}
