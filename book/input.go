package book

import "time"

type CreateBookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	AuthorID    string `json:"authorId" binding:"required"`
	CategoryID  string `json:"categoryId" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
