package controllers

import (
	"ass3_be/helpers"
	"ass3_be/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) Post(c *gin.Context) {
var (
		drink 	models.Drink
		result	gin.H
	)

	c.BindJSON(&drink)
	drink.WaterStatus = helpers.CheckWaterStatus(drink.Water)
	drink.WishStatus = helpers.CheckWaterStatus(drink.Wish)
	idb.DB.Create(&drink)

	result = gin.H{
		"result" : drink,
	}
	
	c.JSON(http.StatusOK, result)
}



func (idb *InDB) All(c *gin.Context) {
	var (
		drink 	[]models.Drink
		result	gin.H
	)

	idb.DB.Order("id desc").Find(&drink)
	
	result = gin.H{
		"result" : drink,
	}
	
	c.JSON(http.StatusOK, result)
}