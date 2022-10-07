package configs

import (
	"ass_02/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB{
	dsn := "root:@tcp(127.0.0.1:3306)/dts_go_ass_2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Order{}, &models.Item{}, )
	return db
}