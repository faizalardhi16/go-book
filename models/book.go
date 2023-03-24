package models

import "time"

type Book struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Title       string
	Description string
	Author      Author
	AuthorID    string
	Category    Category
	CategoryID  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
