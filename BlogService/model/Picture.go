package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Picture struct {
	ID     uuid.UUID `json:"id" gorm:"primaryKey"`
	URL    string    `json:"url" gorm:"type:string"`
	BlogID uuid.UUID `json:"blog_id"`
}

func (picture *Picture) BeforeCreate(scope *gorm.DB) error {
	picture.ID = uuid.New()
	return nil
}
