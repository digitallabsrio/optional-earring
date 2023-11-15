package routes

import (
	controllers "shive-app/controllers"
	"shive-app/middleware"

	"github.com/gin-gonic/gin"
)

func GenreRoutes(router gin.Engine) {
	router.Use(middleware.AuthenticateUser())
	router.POST("/genres/creategenre", controllers.CreateGenre())
	router.GET("/genres/:genre_id", controllers.GetGenre())
	router.GET("/genres", controllers.GetGenres())
	router.PUT("/genres/:genre_id", controllers.EditGenre())
	router.DELETE("/genres/:genre_id", controllers.DeleteGenre())
}
