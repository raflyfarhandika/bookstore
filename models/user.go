package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/raflyfarhandika/bookstore/helper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"username" form:"username" valid:"required~Name required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email required, email~Invalid Email"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password required, minstringlength(6)~Password must be at least 6 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, errCheck := govalidator.ValidateStruct(u)

	if errCheck != nil {
		return errCheck
	}

	u.Password = helper.HashPassword(u.Password)

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	_, errCheck := govalidator.ValidateStruct(u)

	if errCheck != nil {
		return errCheck
	}

	u.Password = helper.HashPassword(u.Password)

	return nil
}
