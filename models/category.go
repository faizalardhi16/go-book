package models

import "time"

type Category struct {
	ID           string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	CategoryName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
