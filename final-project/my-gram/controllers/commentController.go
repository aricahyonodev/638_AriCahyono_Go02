package controllers

import (
	"fmt"
	"my-gram/dto"
	"my-gram/helpers"
	"my-gram/models"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetAllComments(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var (
		comment []models.Comment
		err error
	)

	// Process Get All Data Photo in DB
	err = idb.DB.Model(&models.Comment{}).Preload("Photo").Preload("User").Find(&comment).Error
	if err != nil {
			inRes.ResBadRequest(err.Error())
			return
		}
	
	var commentAllDto []dto.CommentAllDto
	dtoMapper.Map(&commentAllDto, comment)
	inRes.ResStatusOK(commentAllDto)
}

func (idb *InDB) PostComments(c *gin.Context) {
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var (
		user  	models.User
		photo 	models.Photo
		comment models.Comment
		err   	error
	)

	userID   := helpers.GetUserIdJWT(c)

	// Get Request With Raw JSON
	c.BindJSON(&comment)

	// Check User Data in DB
	err = idb.DB.First(&user, userID).Error
	if err != nil {
		inRes.ResStatusNotFound("Your account not found")
		return
	}

	err = idb.DB.First(&photo, comment.PhotoId).Error
	if err != nil {
		inRes.ResStatusNotFound("Your Photo not found")
		return
	}

	comment.PhotoId = photo.PhotoId
	comment.UserId  = user.UserId

	// Process Create Data Comment in DB
	err = idb.DB.Create(&comment).Error
	if err != nil {
		inRes.ResUnprocessableEntity(err)
		return
	}

	var commentCreatedDto dto.CommentCreatedDto
	dtoMapper.Map(&commentCreatedDto, comment)
	inRes.ResStatusCreated(commentCreatedDto)
}

func (idb *InDB) EditComments(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

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
		inRes.ResStatusNotFound("Your comment not found")
		return
	}

	// Get Request With Raw JSON
	c.BindJSON(&comment)

	// Process Update Data Photo in DB
	err 	= idb.DB.Model(&comment).Updates(comment).Error
	if err != nil {
		inRes.ResUnprocessableEntity(err)
		return
	}
	
	var commentEditDto dto.CommentEditDto
	dtoMapper.Map(&commentEditDto, comment)
	inRes.ResStatusOK(commentEditDto)
}

func (idb *InDB) DeleteComments(c *gin.Context){
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var err error

	// Get Params Url
	commentId := c.Param("commentId")

	// Check Comment with Find in DB
	err = idb.DB.First(&models.Comment{}, commentId).Error
	if err != nil {
		inRes.ResStatusNotFound("Your comment not found")
		return
	}

	// Process Delete Comment In DB
	err = idb.DB.Delete(&models.Comment{}, commentId ).Error
	if err != nil {
		inRes.ResBadRequest("Your comment failed deleted")
		return
	}
	
	inRes.ResMsgStatusOK("Your comment has been successfully deleted")
}