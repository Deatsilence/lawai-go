package mongo

import "time"

type Message struct {
	Role        string    `bson:"role"`        // 'user' or 'assistant'
	Content     string    `bson:"content"`     // Content of the message
	CreatedAt   time.Time `bson:"created_at"`  // Message creation time
	TotalTokens int       `bson:"tokens_used"` // Tokens used for the message as summary
}
