package routers

import(
	"belajar-gin/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.POST("/cars", controllers.CreateCar)
	router.PUT("/cars/:carID", controllers.UpdateCar)
	router.GET("/cars/:carID", controllers.GetCar)
	router.GET("/cars", controllers.GetAllCar)
	router.DELETE("/cars/:carID", controllers.DeleteCar)

	return router
}