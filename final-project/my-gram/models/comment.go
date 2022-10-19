package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	CommentId int    `gorm:"primaryKey;autoIncrement" json:"comment_id"`
	Message   string `json:"message" valid:"required~Message wajib terisi"`
	UserId    int    `gorm:"references:UserId" json:"user_id"`
	PhotoId   int    `gorm:"references:PhotoId" json:"photo_id"`
	User      User
	Photo     Photo
	GormModel
}
type CommentValidasi struct {
		Message string `json:"message" valid:"required~Message wajib terisi"`
	}


func (comment *Comment) BeforeCreate(tx *gorm.DB) (err error) {

	var commentValidasi CommentValidasi
	commentValidasi.Message = comment.Message

	_, errCreate := govalidator.ValidateStruct(commentValidasi)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (comment *Comment) BeforeUpdate(tx *gorm.DB) (err error) {

	var commentValidasi CommentValidasi
	commentValidasi.Message = comment.Message

	_, errCreate := govalidator.ValidateStruct(commentValidasi)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}