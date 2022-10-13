package config

import (
	"ass3_be/models"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)
var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "dts_go_ass_3"
	db		 *gorm.DB
	err      error
)

func InitDb() *gorm.DB{
	config  := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn 	:= config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&models.Drink{})
	return db
}

