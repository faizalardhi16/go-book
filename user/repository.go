package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email CheckEmailInput) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email CheckEmailInput) (User, error) {
	var user User

	err := r.db.Raw("SELECT * FROM users WHERE email = ?", email.Email).Scan(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
