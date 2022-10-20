package controllers

import (
	"fmt"
	"go_jwt/database"
	"go_jwt/helpers"
	"go_jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)

func UserRegister(c *gin.Context) {
	db 			:= database.GetDB() 
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{} 

	if contentType == appJSON {
		c.ShouldBind(&User)
	}else{
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "Bad Request",
			"message" : err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id" 		: 		User.ID,
		"email" 	: 		User.Email,
		"full_name" : 		User.FullName,
	})
}

func UserLogin(c *gin.Context)  {
	db  		:= database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ 		 = db, contentType
	User 		:= models.User{}  
	password 	:= ""

	if contentType == appJSON {
		c.ShouldBind(&User)
	}else{
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" 	: "Unauthorized",
			"message" 	: "Invalid email/password",
		})
		return 
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" 	: "Unauthorized",
			"message" 	: "Invalid email/password",
		})
		return 
	}

	token := helpers.GenerateToken(User.ID, User.Email)
	fmt.Println(token)
	fmt.Println("Hello")
	c.JSON(http.StatusOK, gin.H{
		"token" : token,
	})

}