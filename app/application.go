package app

import "github.com/gin-gonic/gin"

/**
Defining the router.
By separating, router to this layer, we can easily change to another router incase if required.
*/
var router = gin.Default()

func StartApplication() {
	// url-handler mapping
	mapUrls()
	router.Run(":8080")
}
