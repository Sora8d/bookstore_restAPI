package access_token

import (
	"bookstoreapi/oauth/utils/errors"
	"strings"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(access_token_id string) (*AccessToken, *errors.RestErr) {
	AccessTokenId := strings.TrimSpace(access_token_id)
	if len(access_token_id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}
	return s.repository.GetById(AccessTokenId)
}
