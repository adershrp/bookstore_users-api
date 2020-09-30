package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/adershrp/bookstore_users-api/domain/user"
	"github.com/adershrp/bookstore_users-api/service"
)

// create user
/**
All handlers should have *gin.Context as parameter
*/
func CreateUser(c *gin.Context) {
	var user user.User
	fmt.Println("Before Parsing", user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("Bytes", bytes)
	if err != nil {
		fmt.Println(err.Error())
		// TODO handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		// TODO handle json parsing error
		return
	}
	createUser, err := service.CreateUser(user)
	if err != nil {
		fmt.Println(err.Error())
		// TODO handle service error
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
