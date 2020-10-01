package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/adershrp/bookstore_users-api/domain/users"
	"github.com/adershrp/bookstore_users-api/service"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

func getUserId(userIdParam string) (int64, *errors.RestError) {
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("couldn't parse the parameter user_id to number")
		return 0, restErr
	}
	return userId, nil
}

// create users
/**
All handlers should have *gin.Context as parameter
*/
func Create(c *gin.Context) {
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
	c.JSON(http.StatusCreated, createUser.Marshall(c.GetHeader("X-Public") == "true"))
}

// get users
func Get(c *gin.Context) {
	userId, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	user, getErr := service.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	/**
	From path variable extract the userid
	*/
	userId, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
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
	c.JSON(http.StatusOK, updateUser.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, restErr := getUserId(c.Param("user_id"))
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	if restErr := service.DeleteUser(userId); restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

/**

 */
func Search(c *gin.Context) {
	status := c.Query("status")
	results, restErr := service.Search(status)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	c.JSON(http.StatusOK, results.Marshall(c.GetHeader("X-Public") == "true"))
}
