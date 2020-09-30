package service

import (
	"github.com/adershrp/bookstore_users-api/domain/user"
	"github.com/adershrp/bookstore_users-api/utils/errors"
)

func CreateUser(user user.User) (*user.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUser(userId int64) (*user.User, *errors.RestError) {
	eUser := &user.User{Id: userId}
	if err := eUser.Get(); err != nil {
		return nil, err
	}
	return eUser, nil
}
