package domain

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string `gorm:"uniqueIndex;not null;size:255" validate:"min=6"`
	PasswordHash string `gorm:"not null" validate:"min=6"`
	Role         string `gorm:"default:'user'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
