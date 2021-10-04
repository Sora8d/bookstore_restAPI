package mysql_utils

import (
	"bookstoreapi/users/utils/errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows       = "no rows in result set"
	errorMailRepeated = "users.email"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error parsing database response"))
	}
	switch sqlErr.Number {
	case 1062:
		if strings.Contains(sqlErr.Message, errorMailRepeated) {
			return errors.NewBadRequestError("email already registered")
		}
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
