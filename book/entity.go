package book

import (
	"go-book/author"
	"go-book/category"
	"time"
)

type Book struct {
	ID          string        `json:"id" binding:"required"`
	Title       string        `json:"title" binding:"required"`
	Description string        `json:"description" binding:"required"`
	AuthorID    string        `json:"authorId" binding:"required"`
	Author      author.Author `json:"author"`
	CategoryID  string        `json:"categoryId" binding:"required"`
	Category    category.Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GetBook struct {
	ID           string
	Title        string
	Description  string
	CategoryName string
	AuthorName   string
	AuthorID     string
	CategoryID   string
}
