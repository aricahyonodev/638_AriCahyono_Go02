package router

import (
	"ass3_be/config"
	"ass3_be/controllers"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func StartApp() *gin.Engine {

	db 	   := config.InitDb()
	inDB   := &controllers.InDB{DB : db}
	r 	   := gin.Default()
	
	r.Use(cors.Default())

	r.GET("/drinks", inDB.All)
	r.POST("/drink", inDB.Post)

	return r
}

