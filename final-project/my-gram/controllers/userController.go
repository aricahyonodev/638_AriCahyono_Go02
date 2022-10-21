package controllers

import (
	"my-gram/dto"
	"my-gram/helpers"
	"my-gram/models"
	"strconv"

	dtoMapper "github.com/dranikpg/dto-mapper"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) Register(c *gin.Context) {

	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	// Get Request With Raw JSON
	var (
		user models.User
		err error
	)
	// Get Request With Form Data
	user.Username 	= c.Request.FormValue("username")
	user.Email 		= c.Request.FormValue("email")
	user.Password 	= c.Request.FormValue("password")
	user.Age, _	    = strconv.Atoi(c.Request.FormValue("age"))
	// Get Request File
	file, err   	:= c.FormFile("profile_image_url")
	filename 		:= ""

	if err == nil {
			filename = helpers.GetFilenamePhoto(file)
		}

	user.ProfileImageUrl = filename

	// Process Create in DB
	err = idb.DB.Create(&user).Error
	if err != nil {
		inRes.ResUnprocessableEntity(err)
		return
	}

	// Map Model to DTO
	var userDto dto.UserDto
	dtoMapper.Map(&userDto, user)
	inRes.ResStatusCreated(userDto)

	// Save Photo in Temp
	c.SaveUploadedFile(file, "temp/profiles/" + filename)
}


func (idb *InDB) Login(c *gin.Context)  {
	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}

	var user models.User
	c.BindJSON(&user)
	password := user.Password

	err := idb.DB.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		inRes.ResStatusUnauthorized("Invalid email/password")
		return 
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))
	if !comparePass{
		inRes.ResStatusUnauthorized("Invalid email/password")
		return 
	}

	token := helpers.GenerateToken(uint(user.UserId), user.Email)
	inRes.ResTokenStatusOK(token)
}
	

func (idb *InDB) Edit(c *gin.Context){

	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}
	
	var (
		user models.User
		err error
	)

	// Get userId With JWT Token
	user.UserId = helpers.GetUserIdJWT(c)
	
	// Get userId With Query Params
	// user.ID , _ = strconv.Atoi(c.Query("userId"))

	// Get Request With Form Data
	user.Username 	= c.Request.FormValue("username")
	user.Email 		= c.Request.FormValue("email")
	// Get Request File
	file, err   	:= c.FormFile("profile_image_url")
	filename 		:= ""

	if err == nil {
			filename = helpers.GetFilenamePhoto(file)
		}

	user.ProfileImageUrl = filename

	// Check Data in DB
	err		   = idb.DB.First(&models.User{}, user.UserId).Error
	if err != nil {
		inRes.ResStatusNotFound("Your account not found")
		return
	}
	
	// Process Update in DB
	err 	   = idb.DB.Model(&user).Updates(user).Error
	if err != nil {
		inRes.ResUnprocessableEntity(err)
		return
	}
	
	// Map Model to DTO
	var userEditDto dto.UserEditDto
	dtoMapper.Map(&userEditDto, user)
	inRes.ResStatusOK(userEditDto)

	// Save Photo in Temp
	c.SaveUploadedFile(file, "temp/profiles/" + filename)
}


func (idb *InDB) Delete(c *gin.Context){

	// Init Response Handler
	inRes := &helpers.InResponse{ C : c}
	
	// Get userId With JWT Token
	userID := helpers.GetUserIdJWT(c)
	
	// Check Data in DB
	var err error
	err		   = idb.DB.First(&models.User{}, userID).Error
	if err != nil {
		inRes.ResStatusNotFound("Your account not found")
		return
	}
	
	//  Process Delete in DB
	err 	   = idb.DB.Delete(&models.User{}, userID).Error
	if err != nil {
		inRes.ResBadRequest("Your account failed deleted")
		return
	}

	inRes.ResMsgStatusOK("Your account has been successfully deleted")
}


// func (idb *InDB) All(c *gin.Context) {
// 	var drink  []models.User
// 	idb.DB.Find(&drink)

// 	c.JSON(http.StatusOK, helpers.ResponseData(200, drink))
// }