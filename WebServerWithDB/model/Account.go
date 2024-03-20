package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	ID    uuid.UUID `json:"id"`
	Username  string    `json:"username" gorm:"not null;type:string"`
	Password string `json:"password" gorm:"not null;type:string"`
	Email string `json:"email" gorm:"not null;type:string"`
	Role Role `json:"role" gorm:"not null;type:string"`
	IsBlocked bool `json:"isblocked" gorm:"not null;type:bool"`
}

func (account *Account) BeforeCreate(scope *gorm.DB) error {
	account.ID = uuid.New()
	return nil
}