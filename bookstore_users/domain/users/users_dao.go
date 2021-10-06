package users

import (
	"bookstoreapi/users/datasources/mysql/users_db"
	"bookstoreapi/users/logger"
	"bookstoreapi/users/utils/date"
	"bookstoreapi/users/utils/errors"
	"bookstoreapi/users/utils/mysql_utils"
	"fmt"
)

const (
	errorNoRows           = "no rows in result set"
	queryInsertUser       = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, status, date_created FROM users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

var usersDB = users_db.Client

func (user *User) Get() *errors.RestErr {
	if err := usersDB.Ping(); err != nil {
		panic(err)
	}
	stmt, err := usersDB.Prepare(queryGetUser)
	if err != nil {
		//This logger we should do with everywhere we have an error
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Status, &user.DateCreated); err != nil {
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
	inserResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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

func (user *User) Update() *errors.RestErr {
	stmt, err := usersDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := usersDB.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestErr) {
	stmt, err := usersDB.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	results := make(Users, 0)
	for rows.Next() {
		var current User
		if err := rows.Scan(&current.Id, &current.FirstName, &current.LastName, &current.Email, &current.DateCreated, &current.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, current)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users matching status %s", status))
	}
	return results, nil
}
