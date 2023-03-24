package category

import "gorm.io/gorm"

type Repository interface {
	CreateCategory(category Category) (Category, error)
	CheckCategoryAvailability(categoryName string) (Category, error)
	FindCategoryId(id string) (Category, error)
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

func (r *repository) CheckCategoryAvailability(categoryName string) (Category, error) {
	var category Category

	err := r.db.Raw("SELECT * from categories where category_name = ?", categoryName).Scan(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) FindCategoryId(id string) (Category, error) {
	var category Category

	err := r.db.Raw("SELECT * from categories where id = ?", id).Scan(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
