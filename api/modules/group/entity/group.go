package entity

import (
	"time"
)

type Group struct {
	ID        string `gorm:"primary_key"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt   time.Time
}

type AccountToGroup struct {
	AccountID string `gorm:"primary_key"`
	GroupID   string `gorm:"primary_key"`
	JoinedAt  time.Time
	LeftAt    time.Time
}
