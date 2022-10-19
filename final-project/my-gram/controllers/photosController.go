package controllers

import (
	"mime/multipart"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetAllPhotos(c *gin.Context){
	var (
		photo []models.Photo
		err error
	)

	// Process Get All Data Photo in DB
	err = idb.DB.Model(&models.Photo{}).Preload("User").Find(&photo).Error
	if err != nil {
			c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, err.Error()))
			return
		}
		
	c.JSON(http.StatusOK,
	helpers.ResponseData(200, photo))
}

func (idb *InDB) PostPhotos(c *gin.Context){

	var (
		user models.User
		photo models.Photo
		err error
	)

	userData   := c.MustGet("userData").(jwt.MapClaims)
	userID 	   := uint(userData["id"].(float64))
	
	// Check Data in DB
	err		   = idb.DB.First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your account not found"))
			return
	}

	// Get Request With Raw JSON
	photo.Title 	= c.Request.FormValue("title")
	photo.Caption 	= c.Request.FormValue("caption")
	// Get Request File
	file, err   	:= c.FormFile("photo")
	filename 		:= ""

	if err == nil {
			filename = getFilenamePhoto(file)
		}

	photo.PhotoUrl = filename
	photo.UserId   = user.UserId

	// Process Create Data Photo in DB
	err = idb.DB.Create(&photo).Error
	if err != nil {
			c.JSON(http.StatusUnprocessableEntity, 
			helpers.ResponseMessage(422, err.Error()))
			return
		}
		
	c.JSON(http.StatusCreated,
	helpers.ResponseData(201, photo))

	// Save Photo in Temp
	c.SaveUploadedFile(file, "temp/" + filename)

}

func (idb *InDB) EditPhotos(c *gin.Context){

	type ResponsePhoto struct {
		Id   	 int   		`json:"id"`
		Title    string 	`json:"title"`
		Caption  string 	`json:"caption"`
		PhotoUrl string 	`json:"photo_url" `
		UserId   int    	`json:"user_url"`
		UpdateAt *time.Time `json:"updated_at,omitempty"`
	}

	var (
		photo models.Photo
		err error
		responsePhoto ResponsePhoto
	)

	// Get Params Url
	photoId := c.Param("photoId")
	
	// Check Data in DB
	err		   = idb.DB.First(&photo, photoId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your photo not found"))
			return
	}

	// Get Request With Raw JSON
	photo.Title 	= c.Request.FormValue("title")
	photo.Caption 	= c.Request.FormValue("caption")
	// Get Request File
	file, err   	:= c.FormFile("photo")
	filename 		:= ""

	if err == nil {
			filename = getFilenamePhoto(file)
		}

	photo.PhotoUrl = filename

	// Process Update Data Photo in DB
	err 	   = idb.DB.Model(&photo).Updates(photo).Error
	if err != nil {
			c.JSON(http.StatusUnprocessableEntity, 
			helpers.ResponseMessage(422, err.Error()))
			return
		}
	
	responsePhoto.Id 		= photo.PhotoId
	responsePhoto.Title 	= photo.Title
	responsePhoto.Caption  	= photo.Caption
	responsePhoto.UpdateAt 	= photo.UpdateAt
	responsePhoto.UserId   	= photo.UserId
	responsePhoto.PhotoUrl  = photo.PhotoUrl

	c.JSON(http.StatusOK,
	helpers.ResponseData(200, responsePhoto))

	// Save Photo in Temp
	c.SaveUploadedFile(file, "temp/" + filename)

}

func (idb *InDB) DeletePhotos(c *gin.Context){
	var err error

	// Get Params Url
	photoId := c.Param("photoId")

	// Check Photo with Find in DB
	err = idb.DB.First(&models.Photo{}, photoId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your photo not found"))
			return
	}

	// Process Delete In DB
	err = idb.DB.Delete(&models.Photo{}, photoId ).Error
	if err != nil {
			c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your photo failed deleted"))
			return
		}

	c.JSON(http.StatusOK,
	helpers.ResponseMessage(200, "Your photo has been successfully deleted"))
}

func getFilenamePhoto(file *multipart.FileHeader ) string{
	// Set Format Save Photo in Temp
	extension := filepath.Ext(file.Filename)
	epoc 	  := strconv.FormatInt(time.Now().Unix(), 10)
	randomFilename  := helpers.RandomString(len(file.Filename))
	return randomFilename + "-" + epoc + extension
}