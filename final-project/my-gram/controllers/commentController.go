package controllers

import (
	"fmt"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetAllComments(c *gin.Context){
	var (
		comment []models.Comment
		err error
	)

	// Process Get All Data Photo in DB
	err = idb.DB.Model(&models.Comment{}).Preload("Photo").Preload("User").Find(&comment).Error
	if err != nil {
			c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, err.Error()))
			return
		}
		
	c.JSON(http.StatusOK,
	helpers.ResponseData(200, comment))
}

func (idb *InDB) PostComments(c *gin.Context) {

	var (
		user  	models.User
		photo 	models.Photo
		comment models.Comment
		err   	error
	)

	userData := c.MustGet("userData").(jwt.MapClaims)
	userID   := uint(userData["id"].(float64))

	// Get Request With Raw JSON
	c.BindJSON(&comment)

	// Check User Data in DB
	err = idb.DB.First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ResponseMessage(400, "Your account not found"))
		return
	}

	err = idb.DB.First(&photo, comment.PhotoId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest,
			helpers.ResponseMessage(400, "Your Photo not found"))
		return
	}

	comment.PhotoId = photo.PhotoId
	comment.UserId  = user.UserId

	// Process Create Data Comment in DB
	err = idb.DB.Create(&comment).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity,
			helpers.ResponseMessage(422, err.Error()))
		return
	}

	commentResult := map[string]interface{}{
	"comment_id": comment.CommentId,
	"message" 	: comment.Message,
	"photo_id" 	: comment.PhotoId,
	"user_id"	: comment.UserId,
	"create_at" : comment.CreateAt,
	}
	
	c.JSON(http.StatusCreated,
		helpers.ResponseData(201, commentResult))
}

func (idb *InDB) EditComments(c *gin.Context){

	var (
		comment models.Comment
		err error
	)

	// Get Params Url
	commentId := c.Param("commentId")
	
	// Check Data in DB
	fmt.Println(commentId)
	err		= idb.DB.First(&comment, commentId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your comment not found"))
			return
	}
	fmt.Println(comment.Message)
	// Get Request With Raw JSON
	c.BindJSON(&comment)
	fmt.Println(comment.Message)

	// Process Update Data Photo in DB
	err 	= idb.DB.Model(&comment).Updates(comment).Error
	if err != nil {
			c.JSON(http.StatusUnprocessableEntity, 
			helpers.ResponseMessage(422, err.Error()))
			return
		}
	
	commentResult := map[string]interface{}{
	"comment_id": comment.CommentId,
	"message" 	: comment.Message,
	"photo_id" 	: comment.PhotoId,
	"user_id"	: comment.UserId,
	"update_at" : comment.UpdateAt,
	}
	
	c.JSON(http.StatusOK,
	helpers.ResponseData(200, commentResult))

}

func (idb *InDB) DeleteComments(c *gin.Context){
	var err error

	// Get Params Url
	commentId := c.Param("commentId")

	// Check Comment with Find in DB
	err = idb.DB.First(&models.Comment{}, commentId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your comment not found"))
			return
	}

	// Process Delete Comment In DB
	err = idb.DB.Delete(&models.Comment{}, commentId ).Error
	if err != nil {
			c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your comment failed deleted"))
			return
		}

	c.JSON(http.StatusOK,
	helpers.ResponseMessage(200, "Your comment has been successfully deleted"))
}