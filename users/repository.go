package users

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByID(userID string) (User, error)
	FindByEmail(email CheckEmailInput) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(userID string) (User, error) {
	var user User

	err := r.db.Where("id = ?", userID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email CheckEmailInput) (User, error) {
	var user User

	err := r.db.Where("email = ?", email.Email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
