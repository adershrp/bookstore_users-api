package users

import (
	"fmt"

	"github.com/adershrp/bookstore_users-api/utils/errors"
)

var dbData = make(map[int64]*User)

func (user *User) Get() *errors.RestError {
	result, ok := dbData[user.Id]
	if !ok {
		return errors.NewNotFoundError(fmt.Sprintf("users %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedBy = result.CreatedBy
	return nil
}

func (user *User) Save() *errors.RestError {
	current := dbData[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registerd", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("users %d already exists", user.Id))
	}
	dbData[user.Id] = user
	return nil
}
