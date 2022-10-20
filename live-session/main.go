package main

import (
	"livecode/middlewares"
	"livecode/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	

	r 	   := gin.Default()
	r.GET("/hello", hello)
	r.Use(middlewares.HelloMiddleware())
	r.Use(middlewares.HelloMiddleware2())
	r.GET("/hello/:hello", hello)
	// r.GET("/hello", hello)
	r.Run(":3000")
}

func hello(c *gin.Context) {
	var biodata models.Biodata

	biodata.NamaLengkap = "Ari Cahyono"
	biodata.KodePeserta = "149368582101-638"
    c.JSON(http.StatusOK, gin.H{
      "message": biodata,
    })
}