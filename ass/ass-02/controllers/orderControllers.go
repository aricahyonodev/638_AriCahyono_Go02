package controllers

import (
	"ass_02/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) Post(c *gin.Context) {
	var (
		order 	models.Order
		result	gin.H
	)

	c.BindJSON(&order)
	idb.DB.Create(&order)
	result = gin.H{
		"result" : order,
	}
	
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) All(c *gin.Context) {
	var (
		orders 	[]models.Order
		result	gin.H
	)

	err := idb.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	_ = err

	result = gin.H{
		"result" : orders,
	}
	
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Delete(c *gin.Context) {
	var (
		order models.Order
		result gin.H
	)

	// Set Order Data
	id 	:= c.Param("id")
	err := idb.DB.First(&order, id).Error
	if err != nil {
		result = gin.H{
			"result" : "data not found",
		}
	}

	// Delete Data In DB
	err = idb.DB.Select("Items").Delete(&order).Error
	if err != nil {
		result = gin.H{
			"result" : "delete failed",
		}
	}else {
		result = gin.H{
			"result" : "data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) Update(c *gin.Context) {
	var (
		order 	models.Order
		result	gin.H
	)

	// Set Order Data
	id, _ := strconv.Atoi(c.Param("orderId"))
	order.OrderId = id
	c.BindJSON(&order)

	// Update Order & Item in Database
	idb.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order)
	idb.DB.Model(&order).Updates(&order)

	result = gin.H{
		"result" : order,
	}
	c.JSON(http.StatusOK, result)
}