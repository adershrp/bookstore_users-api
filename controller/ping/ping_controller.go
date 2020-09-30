package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
Passing the context, its mandatory
*/
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
