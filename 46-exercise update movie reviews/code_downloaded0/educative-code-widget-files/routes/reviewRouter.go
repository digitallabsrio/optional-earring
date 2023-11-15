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
	router.DELETE("reviews/:reviewer_id", controllers.DeleteReview())
	router.GET("reviews/user_reviews/:reviewer_id", controllers.AllUserReviews())
	router.PUT("reviews/user_reviews", controllers.EditReviews())

}
