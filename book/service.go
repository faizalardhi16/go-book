package book

import (
	"fmt"
	"go-book/author"
	"go-book/category"

	"github.com/google/uuid"
)

type Service interface {
	SaveBook(input CreateBookInput) (Book, error)
	GetAllBook() ([]Book, error)
}

type service struct {
	repository         Repository
	authorRepository   author.Repository
	categoryRepository category.Repository
}

func NewBookService(repository Repository, authorRepository author.Repository, categoryRepository category.Repository) *service {
	return &service{repository, authorRepository, categoryRepository}
}

func (s *service) SaveBook(input CreateBookInput) (Book, error) {
	book := Book{}

	newAuthor, err := s.authorRepository.FindAuthorById(input.AuthorID)

	if err != nil {
		return book, err
	}

	newCategory, err := s.categoryRepository.FindCategoryId(input.CategoryID)

	if err != nil {
		return book, err
	}

	book.Author = newAuthor
	book.Category = newCategory
	book.ID = uuid.New().String()
	book.Title = input.Title
	book.Description = input.Description
	book.CategoryID = input.CategoryID
	book.AuthorID = input.AuthorID

	newBook, err := s.repository.CreateBook(book)

	if err != nil {
		return book, err
	}

	return newBook, nil

}

func (s *service) GetAllBook() ([]Book, error) {
	allBook, err := s.repository.GetBook()

	fmt.Println(allBook, "ALL BOOK")
	if err != nil {
		return allBook, err
	}

	return allBook, nil
}
