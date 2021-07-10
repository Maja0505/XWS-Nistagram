package authentication

import (
	"gorm.io/gorm"
	"math/rand"
)

type User struct {
	ID uint64            `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.ID = uint64(rand.Int())
	return nil
}