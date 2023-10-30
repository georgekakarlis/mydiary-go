package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshToken struct {
	gorm.Model
	Token     string
	UserID    uint
	ExpiresAt time.Time
}