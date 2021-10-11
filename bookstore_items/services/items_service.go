package services

import (
	"bookstoreapi/items/domain/items"

	resterrs "github.com/Sora8d/bookstore_utils-go/rest_errors"
)

var ItemsService itemsServiceInterface = &itemsService{}

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, resterrs.RestErr)
	Get(string) (*items.Item, resterrs.RestErr)
}

type itemsService struct{}

func NewService() itemsServiceInterface {
	return &itemsService{}
}

func (s *itemsService) Create(item items.Item) (*items.Item, resterrs.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(string) (*items.Item, resterrs.RestErr) {
	return nil, nil
}
