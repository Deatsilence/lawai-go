package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PasswordReset struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`            // Hem BSON hem JSON için ID desteği
	Email     *string            `bson:"email" json:"email" validate:"required,email"` // Email için BSON, JSON ve validasyon
	Code      string             `bson:"code" json:"code"`                             // Kod için BSON ve JSON
	CreatedAt time.Time          `bson:"createdAt" json:"created_at"`                  // Zaman damgası BSON/JSON uyumlu
	ExpiresAt time.Time          `bson:"expiresAt" json:"expires_at"`                  // Son kullanma zamanı BSON/JSON uyumlu
}
