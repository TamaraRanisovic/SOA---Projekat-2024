package model

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID uuid.UUID `json:"id"`
	CommenterID   string `json:"commenterid"`
	TimePublished time.Time `json:"timepublished"`
	Comment       string `json:"comment"`
	TimeLastEdit time.Time `json:"timelastedit"`
	
}

func (rating *Comment) BeforeCreate(scope *gorm.DB) error {
	rating.ID = uuid.New()
	return nil
}