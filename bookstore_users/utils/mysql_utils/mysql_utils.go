package mysql_utils

import (
	"errors"
	"strings"

	resterrs "github.com/Sora8d/bookstore_utils-go/rest_errors"

	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows       = "no rows in result set"
	errorMailRepeated = "users.email"
)

func ParseError(err error) *resterrs.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			resterr := resterrs.NewNotFoundError("no record matching given id")
			return &resterr
		}
		resterr := resterrs.NewInternalServerError("error parsing database response", errors.New("database error"))
		return &resterr
	}
	switch sqlErr.Number {
	case 1062:
		if strings.Contains(sqlErr.Message, errorMailRepeated) {
			resterr := resterrs.NewBadRequestErr("email already registered")
			return &resterr
		}
		resterr := resterrs.NewBadRequestErr("invalid data")
		return &resterr
	}
	resterr := resterrs.NewInternalServerError("error processing request", errors.New("database error"))
	return &resterr
}
