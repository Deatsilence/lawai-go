package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APILog struct {
	ID               primitive.ObjectID     `bson:"_id,omitempty" json:"id"`                   // Unique log ID
	ChatID           primitive.ObjectID     `bson:"chat_id" json:"chatId"`                     // Chat foreign key
	Request          map[string]interface{} `bson:"request" json:"request"`                    // API request in JSON format
	Response         map[string]interface{} `bson:"response" json:"response"`                  // API response in JSON format
	PromptTokens     int                    `bson:"prompt_tokens" json:"promptTokens"`         // Tokens used for prompt
	CompletionTokens int                    `bson:"completion_tokens" json:"completionTokens"` // Tokens used for completion
	TotalTokens      int                    `bson:"total_tokens" json:"totalTokens"`           // Total tokens used
	CreatedAt        time.Time              `bson:"created_at" json:"createdAt"`               // Log creation time
}
