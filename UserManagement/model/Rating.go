package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Rating struct {
	ID uuid.UUID `json:"id"`
	RatingValue   string `json:"ratingvalue" gorm:"not null;type:int"` //must be in range [1,5]
	Comment       string `json:"comment"`
	TimePublished time.Time `json:"timepublished"`
}

func (rating *Rating) BeforeCreate(scope *gorm.DB) error {
	rating.ID = uuid.New()
	return nil
}