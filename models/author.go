package models

import "time"

type Author struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	AuthorName string `gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
