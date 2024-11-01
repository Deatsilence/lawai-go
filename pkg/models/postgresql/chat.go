package postgresql

import "time"

type Chat struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"index"`          // Owner of the chat
	SystemMessage string    `gorm:"type:text"`      // If role is system, this will be the message
	CreatedAt     time.Time `gorm:"autoCreateTime"` // Set to current time when record is created
}
