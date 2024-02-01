package repository

import (
	"github.com/raflyfarhandika/bookstore/models"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db}
}

func (r *BookRepo) FindAll() ([]models.Book, error) {
	var result []models.Book

	err := r.db.Find(&result).Error

	if err != nil {
		return []models.Book{}, err
	}

	return result, nil
}

func (r *BookRepo) FindOne(id uint) (models.Book, error) {
	var result models.Book

	err := r.db.First(&result, id).Error

	if err != nil {
		return models.Book{}, err
	}

	return result, nil
}

func (r *BookRepo) Create(request models.Book) (models.Book, error) {

	err := r.db.Create(&request).Error

	if err != nil {
		return models.Book{}, err
	}

	return request, nil
}

func (r *BookRepo) Delete(book models.Book) error {

	err := r.db.Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepo) Update(oldBook models.Book, newBook models.Book) (models.Book, error) {

	err := r.db.Model(&oldBook).Updates(&newBook).Error

	if err != nil {
		return models.Book{}, err
	}

	return oldBook, nil
}
