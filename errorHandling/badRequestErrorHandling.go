package errorHandling

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBadRequest(c *gin.Context, err error, msg string) {
	if err != nil {
		c.JSON(http.StatusBadRequest, msg)

		return
	}
}
