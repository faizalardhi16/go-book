package category

import "gorm.io/gorm"

type Repository interface {
	CreateCategory(category Category) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCategory(category Category) (Category, error) {
	err := r.db.Create(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
