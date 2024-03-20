package model

import "time"

type Blog struct {
	Title         string
	Description   string
	DatePublished time.Time
	Pictures string
	Status Status
}