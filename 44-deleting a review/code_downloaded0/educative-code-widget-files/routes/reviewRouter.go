package routes

import (
	controllers "shive-app/controllers"
	"shive-app/middleware"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(router gin.Engine) {
	router.Use(middleware.AuthenticateUser())
	router.POST("reviews/addreview", controllers.AddReview())
	router.GET("reviews/filter", controllers.ViewAMovieReviews())
	router.DELETE("reviews/:review_id", controllers.DeleteReview())
}
