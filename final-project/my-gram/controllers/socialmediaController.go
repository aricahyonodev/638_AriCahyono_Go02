package controllers

import (
	"my-gram/dto"
	"my-gram/helpers"
	"my-gram/models"

	"github.com/dgrijalva/jwt-go"
	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetAllSocialMedia(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}


	var (
		socialMedia models.SocialMedia
		err error
	)

	// Process Get All Data Photo in DB
	err = idb.DB.Model(&models.SocialMedia{}).Preload("User").Find(&socialMedia).Error
	if err != nil {
			inRes.ResBadRequest(err.Error())
			return
		}
	
	var socialmediaAllDto dto.SocialmediaAllDto
	dtoMapper.Map(&socialmediaAllDto, socialMedia)
	inRes.ResStatusOK(socialmediaAllDto)
}

func (idb *InDB) PostSocialMedia(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var (
		user models.User
		socialMedia models.SocialMedia
		err error
	)

	// Get Data in JWT
	userData   := c.MustGet("userData").(jwt.MapClaims)
	userID 	   := uint(userData["id"].(float64))

	// Get Request With Raw JSON
	c.BindJSON(&socialMedia)

		// Check Data User in DB
	err		   = idb.DB.First(&user, userID).Error
	if err != nil {
		inRes.ResStatusNotFound("Your account not found")
		return
	}	

	socialMedia.UserId = user.UserId

	// Check Data Social Media User in DB
	err		   = idb.DB.First(&socialMedia, "user_id = ?", userID).Error
	if err == nil {
		inRes.ResStatusNotFound("Your data social media has been filled")
		return
	}	

	// Process Create Data Social Media in DB
	err = idb.DB.Create(&socialMedia).Error
	if err != nil {
			inRes.ResUnprocessableEntity(err)
			return
		}
	
	var socialmediaCreatedDto dto.SocialmediaCreatedDto
	dtoMapper.Map(&socialmediaCreatedDto, socialMedia)
	inRes.ResStatusCreated(socialmediaCreatedDto)
}

func (idb *InDB) EditSocialMedia(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var (
		socialMedia models.SocialMedia
		socialMediaRequest models.SocialMedia
		err error
	)

	// Get Params Url
	socialMediaId := c.Param("socialMediaId")
	// Get Request With Raw JSON
	c.BindJSON(&socialMediaRequest)

	// Check Data in DB
	err		= idb.DB.First(&socialMedia, socialMediaId).Error
	if err != nil {
		inRes.ResStatusNotFound("Your social media not found")
		return
	}
	
	socialMedia.Name = socialMediaRequest.Name
	socialMedia.SocialMediaUrl = socialMediaRequest.SocialMediaUrl

	// Process Update Data Photo in DB
	err 	= idb.DB.Model(&socialMedia).Updates(socialMedia).Error
	if err != nil {
		inRes.ResUnprocessableEntity(err)
		return
	}
	
	var socialmediaEditDto dto.SocialmediaEditDto
	dtoMapper.Map(&socialmediaEditDto, socialMedia)
	inRes.ResStatusOK(socialmediaEditDto)
}

func (idb *InDB) DeleteSocialMedia(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var err error

	// Get Params Url
	socialMediaId := c.Param("socialMediaId")

	// Check Photo with Find in DB
	err = idb.DB.First(&models.SocialMedia{}, socialMediaId).Error
	if err != nil {
		inRes.ResStatusNotFound("Your social media not found")
		return
	}

	// Process Delete In DB
	err = idb.DB.Delete(&models.SocialMedia{}, socialMediaId).Error
	if err != nil {
		inRes.ResBadRequest("Your social media failed deleted")
		return
	}

	inRes.ResMsgStatusOK("Your social media has been successfully deleted")
}
