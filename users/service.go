package users

import (
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	GetUserByEmail(email CheckEmailInput) (bool, error)
	GetUserByID(userID string) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewUserService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	user := User{}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Role = "user"
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

func (s *service) GetUserByID(userID string) (User, error) {

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return user, err
	}

	return user, nil

}

func (s *service) Login(input LoginInput) (User, error) {
	email := CheckEmailInput{}
	password := input.Password

	email.Email = input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, errors.New("No user found with that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}
