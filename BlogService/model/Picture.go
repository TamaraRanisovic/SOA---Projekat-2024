package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Picture struct {
	ID     uuid.UUID `json:"-" gorm:"primaryKey"`
	URL    string    `json:"url" gorm:"type:string"`
	BlogID uuid.UUID `json:"-"`
}

func (picture *Picture) BeforeCreate(scope *gorm.DB) error {
	picture.ID = uuid.New()
	return nil
}
