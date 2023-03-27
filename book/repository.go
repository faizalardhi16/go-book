package book

import (
	"go-book/author"
	"go-book/category"

	"gorm.io/gorm"
)

type Repository interface {
	CreateBook(book Book) (Book, error)
	GetBook() ([]Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(book Book) (Book, error) {
	err := r.db.Create(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) GetBook() ([]Book, error) {
	var book []Book
	var singleBook Book
	var getBook []GetBook

	err := r.db.Model(&Book{}).Create(map[string]interface{}{
		"id": "books.id",
		"author": map[string]interface{}{
			"authorName": "author_name",
			"id":         "author_id",
		},
		"category": map[string]interface{}{
			"categoryName": "category_name",
			"id":           "author_id",
		},
		"title":       "books.title",
		"description": "books.description",
	}).Select("books.id, author_name, category_name, author_id, category_id, title, description").
		Joins("JOIN authors a ON a.id = books.author_id").
		Joins("JOIN categories c ON c.id = books.category_id").
		Scan(&getBook).Error

	for _, res := range getBook {

		author := author.Author{}
		author.AuthorName = res.AuthorName
		author.ID = res.AuthorID

		category := category.Category{}
		category.CategoryName = res.CategoryName
		category.ID = res.CategoryID

		singleBook.AuthorID = res.AuthorID
		singleBook.CategoryID = res.AuthorID
		singleBook.Author = author
		singleBook.Category = category
		singleBook.Title = res.Title
		singleBook.Description = res.Description
		singleBook.ID = res.ID

		book = append(book, singleBook)

	}

	// err := r.db.Preload("authors").Preload("categories").Find(&book).Error

	// err := r.db.Raw("select * from books b join authors a on a.id = b.author_id join categories c on c.id = category_id").Scan(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil

}
