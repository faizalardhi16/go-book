package author

import (
	"go-book/constant"

	"gorm.io/gorm"
)

type Repository interface {
	CreateAuthor(author Author) (Author, error)
	CheckAuthorAvailability(authorName string) (Author, error)
	DeleteAuthor(id string) (string, error)
	FindAuthorById(id string) (Author, error)
}

type repository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateAuthor(author Author) (Author, error) {
	err := r.db.Create(&author).Error

	if err != nil {
		return author, err
	}

	return author, nil
}

func (r *repository) CheckAuthorAvailability(authorName string) (Author, error) {
	var author Author
	err := r.db.Raw("SELECT * from authors where author_name = ?", authorName).Scan(&author).Error

	if err != nil {
		return author, err
	}

	return author, nil
}

func (r *repository) DeleteAuthor(id string) (string, error) {
	author := Author{}

	err := r.db.Raw("select * from authors where id = ?", id).Scan(&author).Error

	if err != nil {
		return "Failed", err
	}

	if author.ID == "" {
		return constant.ResponseEnum.NotFound, nil
	}

	err = r.db.Raw("delete from authors where id = ?", id).Scan(&author).Error

	if err != nil {
		return "Failed", err
	}

	return "Success", nil
}

func (r *repository) FindAuthorById(id string) (Author, error) {
	var author Author
	err := r.db.Raw("SELECT * from authors where id = ?", id).Scan(&author).Error

	if err != nil {
		return author, err
	}

	return author, nil
}
