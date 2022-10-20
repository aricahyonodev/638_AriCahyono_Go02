package models

import (
	"errors"
	"my-gram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
	dtoMapper "github.com/dranikpg/dto-mapper"
)

type User struct {
	UserId   int 	    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username string		`gorm:"uniqueIndex" json:"username" valid:"required~Username wajib terisi"`
	Email    string		`gorm:"uniqueIndex" json:"email" valid:"required~Email wajib terisi,email~Format Email tidak valid"`
	Password string		`json:"password" valid:"required~Password wajib terisi, minstringlength(6)~Password harus terdiri dari min 6 karakter"`
	Age      int		`json:"age" valid:"required~Age wajib terisi"`
	ProfileImageUrl string	`json:"profile_image_url" valid:"required~Profile Image Url wajib terisi"`
	GormModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if u.Age < 8 {
		err = errors.New("age wajib di atas 8")
		return 
	}

	if errCreate != nil {
		err = errCreate
		return 
	}
	

	err = nil
	u.Password = helpers.HashPass(u.Password)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	type UserValidation struct {
		Username string			`json:"username" valid:"required~Username wajib terisi"`
		Email    string			`json:"email" valid:"required~Email wajib terisi,email~Format Email tidak valid"`
		ProfileImageUrl string	`json:"profile_image_url" valid:"required~Profile Image Url wajib terisi"`
	}

	var userValidation UserValidation 
	dtoMapper.Map(&userValidation, u)
	_, errCreate := govalidator.ValidateStruct(userValidation)

	if errCreate != nil {
		err = errCreate
		return 
	}

	err = nil
	return
}



