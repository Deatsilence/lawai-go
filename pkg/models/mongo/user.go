package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`                            // Primary Key
	Username     *string            `bson:"username,unique" json:"username"`          // Unique username, not null
	Email        *string            `bson:"email,unique" json:"email"`                // Unique email, not null
	Password     *string            `bson:"password" json:"password"`                 // Hashed password, not null
	Token        *string            `bson:"token,unique" json:"token"`                // Authentication token
	RefreshToken *string            `bson:"refresh_token,unique" json:"refreshToken"` // Refresh token
	IsVerified   bool               `bson:"is_verified" json:"isVerified"`            // Email verification status
	CreatedAt    time.Time          `bson:"created_at" json:"createdAt"`              // Auto set to current time
	UpdatedAt    *time.Time         `bson:"updated_at" json:"updatedAt"`
	LastLogin    *time.Time         `bson:"last_login_at" json:"lastLoginAt"` //
	DeletedAt    *time.Time         `bson:"deleted_at" json:"deletedAt"`
}

func (u *User) Map() {
	now, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	u.ID = primitive.NewObjectID()
	u.CreatedAt = now
	u.UpdatedAt = &now
	u.IsVerified = false

}
