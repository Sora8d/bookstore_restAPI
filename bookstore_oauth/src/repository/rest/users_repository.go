package rest

import (
	"bookstoreapi/oauth/utils/errors"
	"bookstoreapi/users/domain/users"
	"encoding/json"
	"log"
	"time"

	rest "github.com/go-resty/resty/v2"
)

type newUser struct {
	Id        int64
	FirstName string
	LastName  string
	Email     string
}

var (
	usersRestClient *rest.Client = rest.New()
	/*{
		HostURL: "http://localhost:8081",
		RetryCount: 4,
	}*/
)

func NewRepository() UsersRepository {
	return &usersRepository{}
}

type UsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func (rR *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}
	usersRestClient.SetTimeout(5 * time.Second)
	usersRestClient.SetHostURL("http://127.0.0.1:8080")
	restreq := usersRestClient.R()
	restreq.SetHeader("Content-Type", "application/json")
	restreq.Method = rest.MethodPost
	restreq.Body = request
	restreq.URL = "/users/login"

	response, err := restreq.Send()
	if err != nil {
		log.Println(err)
		return nil, errors.NewInternalServerError("error in the restclient functionality")
	}
	if response == nil || response.Body() == nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.IsError() {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to log into user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
