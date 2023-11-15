package routes

import (
	controllers "shive-app/controllers"
	"shive-app/middleware"

	"github.com/gin-gonic/gin"
)

func GenreRoutes(router gin.Engine) {
	router.Use(middleware.AuthenticateUser())
	router.POST("/genres/creategenre", controllers.CreateGenre())
}
