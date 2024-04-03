package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null;type:string"`
	Description string    `json:"description" gorm:"type:string"`
	DateCreated time.Time `json:"date_created" gorm:"not null"`
	Pictures    []Picture `json:"pictures,omitempty" gorm:"foreignKey:BlogID;references:ID"`
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	blog.ID = uuid.New()
	return nil
}
