package domain

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Role         string `gorm:"default:'user'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
