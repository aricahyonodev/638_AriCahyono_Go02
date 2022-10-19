package controllers

import (
	"my-gram/helpers"
	"my-gram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetAllSocialMedia(c *gin.Context){
	var (
		socialMedia models.SocialMedia
		err error
	)

	// Process Get All Data Photo in DB
	err = idb.DB.Model(&models.SocialMedia{}).Preload("User").Find(&socialMedia).Error
	if err != nil {
			c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, err.Error()))
			return
		}
		
	c.JSON(http.StatusOK,
	helpers.ResponseData(200, socialMedia))
}

func (idb *InDB) PostSocialMedia(c *gin.Context){

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
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your account not found"))
			return
	}	

	socialMedia.UserId = user.UserId

	// Check Data Social Media User in DB
	err		   = idb.DB.First(&socialMedia, "user_id = ?", userID).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your data social media has been filled"))
			return
	}	

	// Process Create Data Social Media in DB
	err = idb.DB.Create(&socialMedia).Error
	if err != nil {
			c.JSON(http.StatusUnprocessableEntity, 
			helpers.ResponseMessage(422, err.Error()))
			return
		}
		
	c.JSON(http.StatusCreated,
	helpers.ResponseData(201, socialMedia))

}

func (idb *InDB) EditSocialMedia(c *gin.Context){

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
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your social media not found"))
			return
	}
	
	socialMedia.Name = socialMediaRequest.Name
	socialMedia.SocialMediaUrl = socialMediaRequest.SocialMediaUrl

	// Process Update Data Photo in DB
	err 	= idb.DB.Model(&socialMedia).Updates(socialMedia).Error
	if err != nil {
			c.JSON(http.StatusUnprocessableEntity, 
			helpers.ResponseMessage(422, err.Error()))
			return
		}
	
	commentResult := map[string]interface{}{
	"social_media_id"	: socialMedia.SocialMediaId,
	"name" 			 	: socialMedia.Name,
	"social_media_url" 	: socialMedia.SocialMediaUrl,
	"user_id"			: socialMedia.UserId,
	"update_at" 		: socialMedia.UpdateAt,
	}
	
	c.JSON(http.StatusOK,
	helpers.ResponseData(200, commentResult))

}
func (idb *InDB) DeleteSocialMedia(c *gin.Context){
	var err error

	// Get Params Url
	socialMediaId := c.Param("socialMediaId")

	// Check Photo with Find in DB
	err = idb.DB.First(&models.SocialMedia{}, socialMediaId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your social media not found"))
			return
	}

	// Process Delete In DB
	err = idb.DB.Delete(&models.SocialMedia{}, socialMediaId).Error
	if err != nil {
			c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your social media failed deleted"))
			return
		}

	c.JSON(http.StatusOK,
	helpers.ResponseMessage(200, "Your social media has been successfully deleted"))
}
