package service

import (
	"github.com/adershrp/bookstore_users-api/domain/users"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

/**
Create a new user record
*/
func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

/**
Fetch user by Id.
*/
func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

/**
Update User, both partial and full update of the payload
*/
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	/**
	Get the user details based on the path variable passed.
	*/
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if isPartial {
		switch {
		case user.FirstName != "":
			current.FirstName = user.FirstName
		case user.LastName != "":
			current.LastName = user.LastName
		case user.Email != "":
			current.Email = user.Email
		}
	} else {
		/**
		Update the user entity fetched from DB and update all the fields using new data send.
		*/
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Validate(); err != nil {
		return nil, err
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

/**
Delete User by passing userId
*/
func DeleteUser(userId int64) *errors.RestError {
	user := &users.User{Id: userId}
	return user.Delete()
}
