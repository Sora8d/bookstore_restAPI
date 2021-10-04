package services

import (
	"bookstoreapi/users/domain/users"
	"bookstoreapi/users/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	reqUser := users.User{Id: userId}
	if err := reqUser.Get(); err != nil {
		return nil, err
	}
	return &reqUser, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
