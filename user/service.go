package user

import (
	"errors"
	"go-book/constant"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(input RegisterInput) (User, error)
	GetUserByEmail(email CheckEmailInput) (bool, error)
	Login(login LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewServiceUser(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateUser(input RegisterInput) (User, error) {
	user := User{}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email

	user.Role = constant.RoleUser.Customer
	user.ID = uuid.New().String()

	passHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passHash)

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) GetUserByEmail(email CheckEmailInput) (bool, error) {

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == "" {
		return true, nil
	}

	return false, nil

}

func (s *service) Login(login LoginInput) (User, error) {
	email := CheckEmailInput{}

	email.Email = login.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))

	if err != nil {
		return user, err
	}

	return user, nil

}
