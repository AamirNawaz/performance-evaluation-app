package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Users struct {
	ID        *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty""`
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	Password  string              `json:"password"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
}