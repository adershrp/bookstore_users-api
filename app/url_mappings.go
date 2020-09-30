package app

import (
	"github.com/adershrp/bookstore_users-api/controller/ping"
	"github.com/adershrp/bookstore_users-api/controller/users"
)

func mapUrls() {
	// handlers are function's. not invoking them here.
	router.GET("/ping", ping.Ping)
	/*
	   users controller
	*/
	router.GET("/users/:user_id", users.GetUser)
	// router.GET("/users/search", controller.SearchUser)
	router.POST("/users", users.CreateUser)
}
