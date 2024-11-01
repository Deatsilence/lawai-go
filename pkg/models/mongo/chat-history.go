package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChatHistory struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"` // For unique identification of the chat
	UserID    primitive.ObjectID `bson:"user_id"`       // Foreign key for the user
	Messages  []Message          `bson:"messages"`      // Chat history
	CreatedAt time.Time          `bson:"created_at"`    // Chat creation time
}
