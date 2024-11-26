package models

import "time"

type Chat struct {
	ID            uint      `bson:"_id" json:"id"`                       // Primary Key
	UserID        uint      `bson:"user_id,index" json:"userId"`         // Owner of the chat
	SystemMessage *string   `bson:"system_message" json:"systemMessage"` // System message
	CreatedAt     time.Time `bson:"created_at" json:"createdAt"`         // Auto set to current time
}
