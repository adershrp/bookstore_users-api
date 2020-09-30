package service

import (
	"github.com/adershrp/bookstore_users-api/domain/users"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUser(userId int64) (*users.User, *errors.RestError) {
	eUser := &users.User{Id: userId}
	if err := eUser.Get(); err != nil {
		return nil, err
	}
	return eUser, nil
}
