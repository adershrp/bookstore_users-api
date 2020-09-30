package service

import (
	"log"

	"github.com/adershrp/bookstore_users-api/domain/user"
)

func CreateUser(user user.User) (*user.User, error) {
	log.Println("Created User", user)
	return &user, nil
}
