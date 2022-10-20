package middlewares

import (
	"my-gram/configs"
	"my-gram/helpers"
	"my-gram/models"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context){
		// Init Response Handler
		inRes := &helpers.InResponse{ C : c}
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			inRes.ResStatusUnauthorized(err.Error())
			return
		}
		
		c.Set("userData", verifyToken)
		c.Next()
	}
}


func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context){
		// Init Response Handler
		inRes := &helpers.InResponse{ C : c}
		db 	   	 := configs.InitDb()

		// Get JWT
		userID	 := helpers.GetUserIdJWT(c)
		Photo  	 := models.Photo{}
		
		err 	 := db.First(&Photo, userID).Error
		if err != nil {
			inRes.ResStatusNotFound("Photo doesn't exist")
			return
		}

		if Photo.UserId != userID {
			inRes.ResStatusUnauthorized("you are not allowed to access this data")
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context){
		// Init Response Handler
		inRes := &helpers.InResponse{ C : c}
		db 	  := configs.InitDb()

		// Get JWT
		userID	 := helpers.GetUserIdJWT(c)
		Comment  := models.Comment{}
		
		err 	 := db.First(&Comment, userID).Error
		if err != nil {
			inRes.ResStatusNotFound("Comment doesn't exist")
			return
		}

		if Comment.UserId != userID {
			inRes.ResStatusUnauthorized("you are not allowed to access this data")
			return
		}

		c.Next()
	}
}

func SocialMediaAuthorization() gin.HandlerFunc {
	return func(c *gin.Context){
		// Init Response Handler
		inRes 		:= &helpers.InResponse{ C : c}
		db 	  		:= configs.InitDb()

		// Get JWT
		userID	 	:= helpers.GetUserIdJWT(c)
		SocialMedia := models.SocialMedia{}
		
		err 	 := db.First(&SocialMedia, userID).Error
		if err != nil {
			inRes.ResStatusNotFound("Social Media doesn't exist")
			return
		}

		if SocialMedia.UserId != userID {
			inRes.ResStatusUnauthorized("you are not allowed to access this data")
			return
		}

		c.Next()
	}
}