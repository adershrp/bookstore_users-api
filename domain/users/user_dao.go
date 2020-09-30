package users

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"strings"

	"github.com/adershrp/bookstore_users-api/datasources/mysql/users_db"
	"github.com/adershrp/bookstore_users-api/utils/dates"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

const (
	noRowsInResult   = "no rows in result set"
	queryInsertUser  = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUserById = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)
const (
	errorSave           = "error when trying to save user %s"
	errorDuplicateEmail = "email already exists %s"
	errorGetById        = "error when trying to fetch user %d: %s"
	errorUserIdNotFound = "user %d not found"
)

/**
Method Naming convention.
Get - reading data by primary key
all other queries based on other attributes should be called as search
*/
func (user *User) Get() *errors.RestError {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close() // closing the statement. this wil execute before return
	/**
	  stmt.QueryRow returns single row of record hence no need to close the result
	*/
	row := stmt.QueryRow(user.Id)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), noRowsInResult) {
			errors.NewNotFoundError(fmt.Sprintf(errorUserIdNotFound, user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf(errorGetById, user.Id, err.Error()))
	}
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
		sqlError, ok := err.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf(errorSave, err.Error()))
		}
		switch sqlError.Number {
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf(errorDuplicateEmail, user.Email))
		default:
			return errors.NewInternalServerError(fmt.Sprintf(errorSave, err.Error()))
		}
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf(errorSave, err.Error()))
	}
	user.Id = userId
	return nil
}
