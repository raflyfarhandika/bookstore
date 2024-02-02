package repository

import (
	"github.com/raflyfarhandika/bookstore/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) FindOne(id uint) (models.User, error) {
	var result models.User

	err := r.db.First(&result, id).Error

	if err != nil {
		return models.User{}, err
	}

	return result, nil
}

func (r *UserRepo) Create(request models.User) (models.User, error) {

	err := r.db.Create(&request).Error

	if err != nil {
		return models.User{}, err
	}

	return request, nil
}
