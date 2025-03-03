package routing

import (
	"minimalist-calories-api/controllers"
	"minimalist-calories-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupUserRouter(r *gin.Engine) *gin.Engine {
	user := r.Group("/users")
	{
		user.POST("/signup", controllers.SignUp)
		user.POST("/login", controllers.Login)
		user.GET("/validate", middleware.Authorize, controllers.Validate)
		user.POST("/logout", middleware.Authorize, controllers.Logout)
	}

	return r
}
