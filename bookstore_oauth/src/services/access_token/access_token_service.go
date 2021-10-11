package access_token

import (
	"bookstoreapi/oauth/domain/access_token"
	"bookstoreapi/oauth/repository/db"
	"bookstoreapi/oauth/repository/rest"
	"bookstoreapi/oauth/utils/errors"
)

var (
	check_access_token_only map[int]bool = map[int]bool{0: false, 1: true}
	check_at_everything     map[int]bool = map[int]bool{0: true, 1: true}
	grant_type_pass                      = "password"
)

type Repository interface {
	GetById(access_token.AccessToken) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUsersRepo rest.UsersRepository
	dbRepo        db.DbRepository
}

func NewService(dbrepo db.DbRepository, usersrepo rest.UsersRepository) Service {
	return &service{
		restUsersRepo: usersrepo,
		dbRepo:        dbrepo,
	}
}

func (s *service) GetById(access_token_id string) (*access_token.AccessToken, *errors.RestErr) {
	at := access_token.AccessToken{AccessToken: access_token_id}
	if err := at.Validate(check_access_token_only); err != nil {
		return nil, err
	}

	accesToken, err := s.dbRepo.GetById(at)
	if err != nil {
		return nil, err
	}
	return accesToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(check_at_everything); err != nil {
		return nil, err
	}
	//TODO: support both grant types: client credentials and password

	//Athenticate the user against the Users API:
	user, err := s.restUsersRepo.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()

	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}

	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := at.Validate(check_at_everything); err != nil {
		return err
	}
	return s.dbRepo.UpdateExpirationTime(at)
}
