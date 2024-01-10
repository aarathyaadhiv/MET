package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chat struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserIDs   []uint             `json:"user_ids" bson:"user_ids"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	ChatID    uint               `json:"chat_id" bson:"chat_id"`
}


type Message struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SenderID uint               `json:"sender_id" bson:"sender_id"`
	ChatID   uint               `json:"chat_id" bson:"chat_id"`
	Seen     bool               `json:"seen" bson:"seen"`
	// Other fields related to your message content
	MessageContent string `json:"message_content" bson:"message_content"`
	Timestamp      time.Time `json:"timestamp" bson:"timestamp"`
	// Add any additional fields relevant to your messages
}
