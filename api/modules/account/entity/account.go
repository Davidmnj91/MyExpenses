package entity

import (
	entity2 "github.com/Davidmnj91/MyExpenses/modules/group/entity"
	"time"
)

type Account struct {
	ID        string `gorm:"primary_key"`
	Name      string
	Password  Password         `gorm:"embedded;embeddedPrefix:password_"`
	Groups    []*entity2.Group `gorm:"many2many:user_groups;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CloseAt   time.Time
}

type Password struct {
	Hashed string
	Cost   int
}
