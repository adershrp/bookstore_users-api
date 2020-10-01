package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/adershrp/bookstore_users-api/domain/users"
	"github.com/adershrp/bookstore_users-api/service"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

// create users
/**
All handlers should have *gin.Context as parameter
*/
func CreateUser(c *gin.Context) {
	var user users.User
	// there are similar methods for XML, YAML
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	createUser, err := service.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, createUser)
}

// get users
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("couldn't parse the parameter user_id to number")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, getErr := service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// search users
func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

func UpdateUser(c *gin.Context) {
	/**
	From path variable extract the userid
	*/
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("couldn't parse the parameter user_id to number")
		c.JSON(restErr.Status, restErr)
		return
	}
	/**
	Validate the request body
	*/
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	/**
	Set the userId to request payload.
	*/
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	updateUser, serviceErr := service.UpdateUser(isPartial, user)
	if serviceErr != nil {
		c.JSON(serviceErr.Status, serviceErr)
		return
	}
	c.JSON(http.StatusOK, updateUser)
}
