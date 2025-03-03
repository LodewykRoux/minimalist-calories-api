package routing

import (
	"minimalist-calories-api/controllers"
	"minimalist-calories-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupWeightRouter(r *gin.Engine) *gin.Engine {
	weight := r.Group("/weight", middleware.Authorize)
	{
		weight.POST("/save", controllers.SaveWeight)
		weight.DELETE("/delete", controllers.DeleteWeight)
		weight.GET("/getList", controllers.GetWeightList)
	}

	return r
}
