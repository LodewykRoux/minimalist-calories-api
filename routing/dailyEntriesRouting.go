package routing

import (
	"minimalist-calories-api/controllers"
	"minimalist-calories-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupDailyEntriesRouter(r *gin.Engine) *gin.Engine {
	dailyEntry := r.Group("/dailyEntries", middleware.Authorize)
	{
		dailyEntry.POST("/save", controllers.SaveDailyEntry)
		dailyEntry.DELETE("/delete", controllers.DeleteDailyEntry)
		dailyEntry.GET("/getList", controllers.GetDailyEntryList)
	}

	return r
}
