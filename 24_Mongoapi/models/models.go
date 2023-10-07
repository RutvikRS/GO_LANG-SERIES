package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TODOLIST struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Data    string             `json:"data,omitempty"`
	Clicked bool               `json:"clicked,omitempty"`
}
