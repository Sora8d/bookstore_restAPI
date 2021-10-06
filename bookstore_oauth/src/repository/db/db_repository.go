package db

import (
	"bookstoreapi/oauth/clients/cassandra"
	"bookstoreapi/oauth/domain/access_token"
	"bookstoreapi/oauth/utils/errors"

	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_toke, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (dbr *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	// TODO: implement get accesstoken from CassandraDB
	session, err := cassandra.GetSession()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()

	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewBadRequestError("Acess Token Not Found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
