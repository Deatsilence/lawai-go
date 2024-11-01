package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APILog struct {
	ID               primitive.ObjectID     `bson:"_id,omitempty"`     // Log kaydı için ObjectID
	ChatID           primitive.ObjectID     `bson:"chat_id"`           // Hangi chat'e ait olduğu
	Request          map[string]interface{} `bson:"request"`           // Yapılan API isteği (JSON formatında)
	Response         map[string]interface{} `bson:"response"`          // Alınan API yanıtı (JSON formatında)
	PromptTokens     int                    `bson:"prompt_tokens"`     // Prompt için kullanılan token sayısı
	CompletionTokens int                    `bson:"completion_tokens"` // Yanıt için kullanılan token sayısı
	TotalTokens      int                    `bson:"total_tokens"`      // Toplam kullanılan token sayısı
	CreatedAt        time.Time              `bson:"created_at"`        // Log kaydının zamanı
}
