package routing

import (
	"minimalist-calories-api/controllers"
	"minimalist-calories-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupFoodRouter(r *gin.Engine) *gin.Engine {

	food := r.Group("/food", middleware.Authorize)
	{
		food.POST("/save", controllers.SaveFood)
		food.DELETE("/delete", controllers.DeleteFood)
		food.GET("/getList", controllers.GetFoodList)
	}

	return r
}
