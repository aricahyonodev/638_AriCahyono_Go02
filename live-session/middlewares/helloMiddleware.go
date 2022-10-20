package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("hello")
		fmt.Printf("Hello Middleware satu %s \n", name)
		c.Next()
	}
}
func HelloMiddleware2() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Param("hello")
		fmt.Printf("Hello Middleware dua %s \n", name)
		c.Next()
	}
}