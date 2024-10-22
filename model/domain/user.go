package domain

import (
	"time"

)

// User represents a user model for the database
type User struct {
    ID       uint   `gorm:"primaryKey"`                // ID field will automatically be the primary key
    Username string `gorm:"unique;not null"`           // Username must be unique and not null
    Password string `gorm:"not null"`                  // Password must not be null
    CreatedAt time.Time `gorm:"autoCreateTime"`           // Tracks the creation time
    UpdatedAt time.Time `gorm:"autoUpdateTime"`           // Tracks the last update time
}
