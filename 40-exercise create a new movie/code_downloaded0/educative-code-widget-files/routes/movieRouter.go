package routes

import (
	controllers "shive-app/controllers"
	"shive-app/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(router gin.Engine) {
	router.Use(middleware.AuthenticateUser())
	router.POST("/movies/createmovie", controllers.CreateMovie())
}
