package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHistory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`     // Unique identification
	UserID    primitive.ObjectID `bson:"user_id" json:"userId"`       // User foreign key
	Messages  []Message          `bson:"messages" json:"messages"`    // Chat history messages
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"` // Creation time
}
