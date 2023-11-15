package routes

import (
	controllers "shive-app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("users/signup", controllers.CreateUser())
}
