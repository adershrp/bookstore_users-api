package app

import (
	"github.com/adershrp/bookstore_users-api/controller/ping"
	"github.com/adershrp/bookstore_users-api/controller/user"
)

func mapUrls() {
	// handlers are function's. not invoking them here.
	router.GET("/ping", ping.Ping)
	/*
	   user controller
	*/
	router.GET("/users/:user_id", user.GetUser)
	// router.GET("/users/search", controller.SearchUser)
	router.POST("/users", user.CreateUser)
}
