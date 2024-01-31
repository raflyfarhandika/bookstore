package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Page        string `json:"page"`
	Author      string `json:"author"`
	Year        string `json:"year"`
	Image       string `json:"image"`
	Price       uint64 `json:"price"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	_, errCheck := govalidator.ValidateStruct(b)

	if errCheck != nil {
		return errCheck
	}

	return nil
}

func (b *Book) BeforeUpdate(tx *gorm.DB) error {
	_, errCheck := govalidator.ValidateStruct(b)

	if errCheck != nil {
		return errCheck
	}

	return nil
}
