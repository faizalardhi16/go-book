package category

import "github.com/google/uuid"

type Service interface {
	SaveCategory(input CategoryInput) (Category, error)
}

type service struct {
	repository Repository
}

func NewCategoryService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SaveCategory(input CategoryInput) (Category, error) {
	category := Category{}
	category.ID = uuid.New().String()
	category.CategoryName = input.CategoryName

	categories, err := s.repository.CreateCategory(category)

	if err != nil {
		return categories, err
	}

	return categories, nil
}
