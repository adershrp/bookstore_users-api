package users

import (
	"fmt"
	"strings"

	"github.com/adershrp/bookstore_users-api/datasources/mysql/users_db"
	"github.com/adershrp/bookstore_users-api/utils/dates"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_unique"
	queryInsertUser  = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
)
const (
	errorSave           = "error when trying to save user %s"
	errorDuplicateEmail = "email already exists %s"
)

var dbData = make(map[int64]*User)

func (user *User) Get() *errors.RestError {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
	result, ok := dbData[user.Id]
	if !ok {
		return errors.NewNotFoundError(fmt.Sprintf("users %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {
	/**
	  Prepare Statement - has lot of advantage compare to statements.
	  1. Before executing, can validate the sql statement. So if there is a error, can skip db call.
	*/
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // closing the statement. this wil execute before return

	// assigning current system date
	user.DateCreated = dates.GetNowString()
	result, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	/**
	  Instead of using Prepare Statement, we could achieve the same result using below code.
	  result, err = users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	*/

	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {

			return errors.NewBadRequestError(fmt.Sprintf(errorDuplicateEmail, user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf(errorSave, err.Error()))
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf(errorSave, err.Error()))
	}
	user.Id = userId
	return nil
}
