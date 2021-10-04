package users

import (
	"bookstoreapi/users/datasources/mysql/users_db"
	"bookstoreapi/users/utils/date"
	"bookstoreapi/users/utils/errors"
	"bookstoreapi/users/utils/mysql_utils"
)

const (
	errorNoRows     = "no rows in result set"
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id = ?"
)

var usersDB = users_db.Client

func (user *User) Get() *errors.RestErr {
	if err := usersDB.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersDB.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

/*
func (user *User) Get(userId int64) (*User, *errors.RestErr) {
	return nil, nil
}
*/
func (user *User) Save() *errors.RestErr {
	stmt, err := usersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = date.GetNowString()
	inserResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, err := inserResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId

	return nil
}
