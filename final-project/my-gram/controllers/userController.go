package controllers

import (
	"my-gram/helpers"
	"my-gram/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) Register(c *gin.Context) {
	var (
		user 	models.User
		result	gin.H
	)

	// Get Request With Raw JSON
	c.BindJSON(&user)

	// Process Create in DB
	err := idb.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, 
			helpers.ResponseMessage(422, err.Error()))
			return
	}

	userResult := map[string]interface{}{
	"age" 		: user.Age,
	"email" 	: user.Email,
	"id" 		: 1,
	"username" 	: user.Username,
	}

	result = gin.H{
		"status" : 201,
		"result" : userResult,
	}
	
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Login(c *gin.Context)  {

	var (
		user 	models.User
	)

	c.BindJSON(&user)
	password := user.Password

	err := idb.DB.Debug().Where("email = ?", user.Email).Take(&user).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" 	: "Unauthorized",
			"message" 	: "Invalid email/password",
		})
		return 
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error" 	: "Unauthorized",
			"message" 	: "Invalid email/password",
		})
		return 
	}

	token := helpers.GenerateToken(uint(user.UserId), user.Email)
	c.JSON(http.StatusOK, gin.H{
		"token" : token,
	})
}
	
func (idb *InDB) Edit(c *gin.Context){

	var (
		user models.User
		err error
	)

	// Get userId With JWT Token
	userData   	:= c.MustGet("userData").(jwt.MapClaims)
	user.UserId	= int(userData["id"].(float64))
	
	// Get userId With Query Params
	// user.ID , _ = strconv.Atoi(c.Query("userId"))

	// Get Request With Raw JSON
	c.BindJSON(&user)

	// Check Data in DB
	err		   = idb.DB.First(&models.User{}, user.UserId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your account not found"))
			return
	}
	
	// Process Update in DB
	err 	   = idb.DB.Model(&user).Updates(user).Error
	if err != nil {
			c.JSON(http.StatusUnprocessableEntity, 
			helpers.ResponseMessage(422, err.Error()))
			return
		}
		
	c.JSON(http.StatusOK,
	helpers.ResponseData(200, user))
}


func (idb *InDB) Delete(c *gin.Context){

	var (
		err error
	)
	userData   := c.MustGet("userData").(jwt.MapClaims)
	userID 	   := uint(userData["id"].(float64))
	
	// Check Data in DB
	err		   = idb.DB.First(&models.User{}, userID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your account not found"))
			return
	}
	
	//  Process Delete in DB
	err 	   = idb.DB.Delete(&models.User{}, userID).Error
	if err != nil {
			c.JSON(http.StatusBadRequest, 
			helpers.ResponseMessage(400, "Your account failed deleted"))
			return
		}

	c.JSON(http.StatusBadRequest,
	helpers.ResponseMessage(200, "Your account has been successfully deleted"))
}


func (idb *InDB) All(c *gin.Context) {
	var (
		drink  []models.User
		result gin.H
	)

	idb.DB.Find(&drink)

	result = gin.H{
		"result": drink,
	}

	c.JSON(http.StatusOK, result)
}