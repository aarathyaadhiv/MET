package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Users     []uint             `json:"users" bson:"users"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type Message struct {
	Message string `json:"message"`
}


