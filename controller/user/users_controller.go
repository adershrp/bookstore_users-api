package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/adershrp/bookstore_users-api/domain/user"
	"github.com/adershrp/bookstore_users-api/service"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

// create user
/**
All handlers should have *gin.Context as parameter
*/
func CreateUser(c *gin.Context) {
	var nUser user.User
	// there are similar methods for XML, YAML
	if err := c.ShouldBindJSON(&nUser); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	createUser, err := service.CreateUser(nUser)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, createUser)
}

// get user
func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("couldn't parse the parameter user_id to number")
		c.JSON(restErr.Status, restErr)
		return
	}
	eUser, getErr := service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, eUser)
}

// search user
func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
