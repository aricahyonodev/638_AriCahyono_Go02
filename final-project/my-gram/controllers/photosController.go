package controllers

import (
	"my-gram/dto"
	"my-gram/helpers"
	"my-gram/models"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetAllPhotos(c *gin.Context){

	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}
	
	var (
		photo []models.Photo
		err error
	)

	// Process Get All Data Photo in DB
	err = idb.DB.Model(&models.Photo{}).Preload("User").Find(&photo).Error
	if err != nil {
			inRes.ResBadRequest(err.Error())
			return
		}
	
	
	var photoAllDto []dto.PhotoAllDto
	dtoMapper.Map(&photoAllDto, photo)
	inRes.ResStatusOK(photoAllDto)
}


func (idb *InDB) PostPhotos(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var (
		user models.User
		photo models.Photo
		err error
	)

	userID 	   := helpers.GetUserIdJWT(c)
	
	// Check Data in DB
	err		   = idb.DB.First(&user, userID).Error
	if err != nil {
		inRes.ResStatusNotFound("Your account not found")
		return
	}

	// Get Request With Raw JSON
	photo.Title 	= c.Request.FormValue("title")
	photo.Caption 	= c.Request.FormValue("caption")
	// Get Request File
	file, err   	:= c.FormFile("photo")
	filename 		:= ""

	if err == nil {
			filename = helpers.GetFilenamePhoto(file)
		}

	photo.PhotoUrl = filename
	photo.UserId   = user.UserId

	// Process Create Data Photo in DB
	err = idb.DB.Create(&photo).Error
	if err != nil {
			inRes.ResUnprocessableEntity(err)
			return
		}

	var photoCreatedDto dto.PhotoCreateDto
	dtoMapper.Map(&photoCreatedDto, photo)
	inRes.ResStatusCreated(photoCreatedDto)

	// Save Photo in Temp
	c.SaveUploadedFile(file, "temp/photos/" + filename)

}

func (idb *InDB) EditPhotos(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var (
		photo models.Photo
		err error
	)

	// Get Params Url
	photoId := c.Param("photoId")
	
	// Check Data in DB
	err		   = idb.DB.First(&photo, photoId).Error
	if err != nil {
		inRes.ResStatusNotFound("Your photo not found")
		return
	}

	// Get Request With Raw JSON
	photo.Title 	= c.Request.FormValue("title")
	photo.Caption 	= c.Request.FormValue("caption")
	// Get Request File
	file, err   	:= c.FormFile("photo")
	filename 		:= ""

	if err == nil {
		filename = helpers.GetFilenamePhoto(file)
	}

	photo.PhotoUrl = filename

	// Process Update Data Photo in DB
	err 	   = idb.DB.Model(&photo).Updates(photo).Error
	if err != nil {
			inRes.ResUnprocessableEntity(err)
			return
		}
	
	var photoEditDto dto.PhotoEditDto
	dtoMapper.Map(&photoEditDto, photo)
	inRes.ResStatusOK(photoEditDto)

	// Save Photo in Temp
	c.SaveUploadedFile(file, "temp/photos/" + filename)

}

func (idb *InDB) DeletePhotos(c *gin.Context){

	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	// Get Params Url
	photoId := c.Param("photoId")

	// Check Photo with Find in DB
	var err error
	err = idb.DB.First(&models.Photo{}, photoId).Error
	if err != nil {
		inRes.ResStatusNotFound("Your photo not found")
		return
	}

	// Process Delete In DB
	err = idb.DB.Delete(&models.Photo{}, photoId ).Error
	if err != nil {
			inRes.ResBadRequest("Your photo failed deleted")
			return
		}

	inRes.ResMsgStatusOK("Your photo has been successfully deleted")
}

