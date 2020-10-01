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
	router.GET("/users/:user_id", users.Get)
	// router.GET("/users/search", controller.SearchUser)

	/**
	POST - create, not Idempotent, Create multiple record for each operation
	Adding resource under the collection, here add a new user under users collection.
	*/
	router.POST("/users", users.Create)

	/**
	PUT - for UPDATE, Idempotent, PUT replaces the operation entirely
	If the requested URI already exists, will update it. Else create one.
	Idempotent. Same operation will only update (entire record).
	*/
	router.PUT("/users/:user_id", users.Update)
	/**
	PATCH - for partial update, Idempotent, update selected fields
	*/
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
}
