package user

import (
	"net/http"

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
	var user user.User
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

// get user
func GetUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// search user
func SearchUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
