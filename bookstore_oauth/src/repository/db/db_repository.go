package db

import (
	"bookstoreapi/oauth/clients/cassandra"
	"bookstoreapi/oauth/domain/access_token"
	"bookstoreapi/oauth/utils/errors"

	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(access_token.AccessToken) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (dbr *dbRepository) GetById(at access_token.AccessToken) (*access_token.AccessToken, *errors.RestErr) {
	// TODO: implement get accesstoken from CassandraDB
	session := cassandra.GetSession()
	if err := session.Query(queryGetAccessToken, at.AccessToken).Scan(
		&at.AccessToken,
		&at.UserId,
		&at.ClientId,
		&at.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewBadRequestError("Access Token Not Found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &at, nil
}

func (dbr *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()
	if err := session.Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (dbr *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()
	if err := session.Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken,
	).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func NewRepository() DbRepository {
	return &dbRepository{}
}
