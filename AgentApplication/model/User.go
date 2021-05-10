package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID  uuid.UUID `json:"id"`
	FirstName string `json:"firstName" gorm:"not null"`
	LastName string `json:"lastName" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uuid.New()
	return nil
}