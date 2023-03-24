package book

import (
	"go-book/author"
	"go-book/category"
)

type ResponseBook struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Author      author.Author `json:"author"`
	Category    category.Category
}

type ResponseAuthor struct {
	ID         string `json:"id"`
	AuthorName string `json:"authorName"`
}

type ResponseCategory struct {
	ID           string `json:"id"`
	CategoryName string `json:"categoryName"`
}

type ResponseAllBook struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Author      ResponseAuthor   `json:"author"`
	Category    ResponseCategory `json:"category"`
}

func FormatResponseBook(book Book) ResponseBook {
	res := ResponseBook{}

	res.ID = book.ID
	res.Title = book.Title
	res.Description = book.Description
	res.Author = book.Author
	res.Category = book.Category

	return res
}

func FormatGetAllBook(book []Book) []ResponseAllBook {
	res := ResponseAllBook{}
	responseMapper := []ResponseAllBook{}

	for _, to := range book {
		author := ResponseAuthor{}
		author.AuthorName = to.Author.AuthorName
		author.ID = to.AuthorID

		category := ResponseCategory{}
		category.CategoryName = to.Category.CategoryName
		category.ID = to.Category.ID

		res.Author = author
		res.Category = category
		res.ID = to.ID
		res.Description = to.Description
		res.Title = to.Title

		responseMapper = append(responseMapper, res)
	}

	return responseMapper

}
