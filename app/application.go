package app

import (
	"github.com/adershrp/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

/**
Defining the router.
By separating, router to this layer, we can easily change to another router incase if required.
*/
var router = gin.Default()

func StartApplication() {
	// url-handler mapping
	mapUrls()
	logger.Info("Server started on port 9090")
	router.Run(":9090")
}
