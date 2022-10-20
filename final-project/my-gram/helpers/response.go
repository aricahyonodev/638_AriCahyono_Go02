package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseMessage(status int, result interface{}) gin.H {
	return gin.H{
		"status" : status,
		"message": result,
	}
}
func ResponseData(status int, result interface{}) gin.H {
	return gin.H{
		"status": status,
		"data"	: result,
	}
}

type InResponse struct {
	C *gin.Context
}




// 200
func (inRes *InResponse) ResStatusOK(data interface{})  {
		inRes.C.JSON(http.StatusOK, 
			ResponseData(http.StatusOK, data))
}

func (inRes *InResponse) ResMsgStatusOK(msg string)  {
		inRes.C.JSON(http.StatusOK, 
			ResponseMessage(http.StatusOK, msg))
}

func (inRes *InResponse) ResTokenStatusOK(token string)  {
		inRes.C.JSON(http.StatusOK, gin.H{
		"token" : token,
	})
}
// 201
func (inRes *InResponse) ResStatusCreated(data interface{})  {
		inRes.C.JSON(http.StatusCreated, 
			ResponseData(http.StatusCreated, data ))
}

// 400
func (inRes *InResponse) ResBadRequest(msg interface{})  {
		inRes.C.JSON(http.StatusBadRequest, 
			ResponseMessage(http.StatusBadRequest, msg))
}

// 401
func (inRes *InResponse) ResStatusUnauthorized(msg string)  {
		inRes.C.JSON(http.StatusUnauthorized, gin.H{
			"error" 	: "Unauthorized",
			"message" 	: msg,
		})
}

// 404
func (inRes *InResponse) ResStatusNotFound(msg string)  {
		inRes.C.JSON(http.StatusNotFound, 
			ResponseMessage(http.StatusBadRequest, msg))
}

// 422
func (inRes *InResponse) ResUnprocessableEntity(err error) {
		inRes.C.JSON(http.StatusUnprocessableEntity,
			ResponseMessage(http.StatusUnprocessableEntity, err.Error()))
}