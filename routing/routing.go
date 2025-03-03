package routing

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	SetupUserRouter(r)

	SetupFoodRouter(r)

	SetupDailyEntriesRouter(r)

	SetupWeightRouter(r)

	return r
}
