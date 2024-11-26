package models

import "time"

type Message struct {
	Role        *string   `bson:"role" json:"role"`               // 'user' or 'assistant'
	Content     *string   `bson:"content" json:"content"`         // Message content
	CreatedAt   time.Time `bson:"created_at" json:"createdAt"`    // Creation time
	TotalTokens int       `bson:"tokens_used" json:"totalTokens"` // Summary of tokens used
}
