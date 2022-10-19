package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	SocialMediaId  int    `gorm:"primaryKey;autoIncrement" json:"Social_media_id"`
	Name           string `json:"name" valid:"required~Name wajib terisi"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~SocialMediaUrl wajib terisi"`
	UserId         int	   `gorm:"references:user_id"`
	User		   User		
	GormModel
}


type SocialMediaValidation struct {
	Name           string `json:"name" valid:"required~Name wajib terisi"`
	SocialMediaUrl string `json:"social_media_url" valid:"required~SocialMediaUrl wajib terisi"`
}

func (socialMedia *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	
	var smValidation SocialMediaValidation
	smValidation.Name = socialMedia.Name
	smValidation.SocialMediaUrl = socialMedia.SocialMediaUrl

	_, errCreate := govalidator.ValidateStruct(smValidation)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}


func (socialMedia *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	
	var smValidation SocialMediaValidation
	smValidation.Name = socialMedia.Name
	smValidation.SocialMediaUrl = socialMedia.SocialMediaUrl

	_, errCreate := govalidator.ValidateStruct(smValidation)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}