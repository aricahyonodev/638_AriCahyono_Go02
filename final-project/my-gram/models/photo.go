package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	PhotoId  int     `gorm:"primaryKey;autoIncrement" json:"photo_id"`
	Title    string `json:"title" valid:"required~Title wajib terisi"`
	Caption  string `json:"caption" valid:"required~Caption wajib terisi"`
	PhotoUrl string `json:"photo_url" valid:"required~Photo Url wajib terisi"`
	UserId   int    `gorm:"references:user_id"`
	User	 User
	GormModel
}

type PhotoValiation struct {
		Title    string `json:"title" valid:"required~Title wajib terisi"`
		Caption  string `json:"caption" valid:"required~Caption wajib terisi"`
		PhotoUrl string `json:"photo_url" valid:"required~Photo Url wajib terisi"`
	}

func (photo *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	
	var newPhoto PhotoValiation
	newPhoto.Title    = photo.Title
	newPhoto.Caption  = photo.Caption
	newPhoto.PhotoUrl = photo.PhotoUrl

	_, errCreate := govalidator.ValidateStruct(newPhoto)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}


func (photo *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	
	var newPhoto PhotoValiation
	newPhoto.Title    = photo.Title
	newPhoto.Caption  = photo.Caption
	newPhoto.PhotoUrl = photo.PhotoUrl
	_, errCreate := govalidator.ValidateStruct(newPhoto)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}