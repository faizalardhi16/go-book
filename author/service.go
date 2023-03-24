package author

import (
	"go-book/constant"

	"github.com/google/uuid"
)

type Service interface {
	SaveAuthor(authorInput AuthorInput) (Author, error)
	CheckAuthor(authorName string) (bool, error)
	DestroyAuthor(deleteInput DeleteAuthorInput) (string, error)
}

type service struct {
	repository Repository
}

func NewAuthorService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveAuthor(input AuthorInput) (Author, error) {
	author := Author{}

	author.AuthorName = input.AuthorName
	author.ID = uuid.New().String()

	newAuthor, err := s.repository.CreateAuthor(author)

	if err != nil {
		return newAuthor, err
	}

	return newAuthor, nil
}

func (s *service) CheckAuthor(authorName string) (bool, error) {
	author, err := s.repository.CheckAuthorAvailability(authorName)

	if err != nil {
		return false, err
	}

	if author.ID == "" {
		return true, nil
	}

	return false, nil
}

func (s *service) DestroyAuthor(deleteInput DeleteAuthorInput) (string, error) {

	authorResponse, err := s.repository.DeleteAuthor(deleteInput.ID)

	if err != nil {
		return constant.ResponseEnum.Failed, err
	}

	if authorResponse == constant.ResponseEnum.NotFound {
		return constant.ResponseEnum.NotFound, nil
	}

	return constant.ResponseEnum.Success, nil
}
