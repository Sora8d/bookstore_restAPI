//Later this will be made into a shared library
package localerrors

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string) error {
	return errors.New(message)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

//This is from the lesson 39/48, wirdly, this is not on the utils

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: "unable to retrieve user information from given access_token",
		Status:  http.StatusUnauthorized,
		Error:   "unauthorized",
	}
}
