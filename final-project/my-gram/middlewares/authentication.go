package middlewares

import (
	"my-gram/configs"
	"my-gram/helpers"
	"my-gram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context){
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" 	: "Unauthenticated",
				"message" 	: err.Error(),
			}) 
			return
		}
		
		c.Set("userData", verifyToken)
		c.Next()
	}
}


func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context){
		db 	   	 := configs.InitDb()

		// Get JWT
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID	 := int(userData["id"].(float64))
		Photo  	 := models.Photo{}
		
		err 	 := db.First(&Photo, userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error"   : "Data Not Found",
				"message" : "data doesn't exist",
			})
			return
		}

		if Photo.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error"   : "unauthorized",
				"message" : "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context){
		db 	   	 := configs.InitDb()

		// Get JWT
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID	 := int(userData["id"].(float64))
		Comment  	 := models.Comment{}
		
		err 	 := db.First(&Comment, userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error"   : "Data Not Found",
				"message" : "data doesn't exist",
			})
			return
		}

		if Comment.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error"   : "unauthorized",
				"message" : "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context){
		db 	   	 := configs.InitDb()

		// Get JWT
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID	 := int(userData["id"].(float64))
		SocialMedia  	 := models.SocialMedia{}
		
		err 	 := db.First(&SocialMedia, userID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error"   : "Data Not Found",
				"message" : "data doesn't exist",
			})
			return
		}

		if SocialMedia.UserId != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error"   : "unauthorized",
				"message" : "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}