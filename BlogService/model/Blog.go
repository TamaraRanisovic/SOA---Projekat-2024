package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blog struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null;type:string"`
	Description string    `json:"description" gorm:"not null;type:string"`
	DateCreated time.Time `json:"date_created" gorm:"not null;type:timestamp"`
	Status      Status    `json:"status" gorm:"not null;type:string"`
	Pictures    []Picture `json:"pictures,omitempty" gorm:"foreignKey:BlogID;references:ID"`
}

func (blog *Blog) BeforeCreate(scope *gorm.DB) error {
	blog.ID = uuid.New()
	return nil
}

func ValidateJsonBlog(blog Blog) error {
	if blog.Title == "" {
		return errors.New("title cannot be empty")
	}

	if blog.Description == "" {
		return errors.New("description cannot be empty")
	}

	if blog.DateCreated.IsZero() {
		return errors.New("date cannot be empty")
	}
	return nil
}
