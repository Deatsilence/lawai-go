package postgresql

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`       // Hashed password
	CreatedAt time.Time `gorm:"autoCreateTime"` // Set to current time when record is created
}
