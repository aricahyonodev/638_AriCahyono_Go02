package router

import (
	"github.com/gin-gonic/gin"
	"my-gram/configs"
	"my-gram/controllers"
	"my-gram/middlewares"
)

func StartApp() *gin.Engine {

	db 	   := configs.InitDb()
	inDB   := &controllers.InDB{DB : db}
	r 	   := gin.Default()
	
	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", inDB.Register)
		userRouter.POST("/login", inDB.Login)
		userRouter.Use(middlewares.Authentication())
		userRouter.DELETE("/", inDB.Delete)
		userRouter.PUT("/", inDB.Edit)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", inDB.PostPhotos)
		photoRouter.GET("/", inDB.GetAllPhotos)

		photoRouter.Use(middlewares.PhotoAuthorization())
		photoRouter.DELETE("/:photoId", inDB.DeletePhotos)
		photoRouter.PUT("/:photoId", inDB.EditPhotos)
	}

	commentsRouter := r.Group("/comments")
	{
		commentsRouter.Use(middlewares.Authentication())
		commentsRouter.POST("/", inDB.PostComments)
		commentsRouter.GET("/", inDB.GetAllComments)

		commentsRouter.Use(middlewares.CommentAuthorization())
		commentsRouter.DELETE("/:commentId", inDB.DeleteComments)
		commentsRouter.PUT("/:commentId", inDB.EditComments)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", inDB.PostSocialMedia)
		socialMediaRouter.GET("/", inDB.GetAllSocialMedia)

		socialMediaRouter.Use(middlewares.SocialMediaAuthorization())
		socialMediaRouter.DELETE("/:socialMediaId", inDB.DeleteSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", inDB.EditSocialMedia)
	}

	return r
}

