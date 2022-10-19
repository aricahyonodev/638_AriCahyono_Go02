package helpers

import "github.com/gin-gonic/gin"

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