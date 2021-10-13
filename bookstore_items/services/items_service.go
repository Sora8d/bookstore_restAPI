package services

import (
	"bookstoreapi/items/domain/items"
	"bookstoreapi/items/domain/queries"

	resterrs "github.com/Sora8d/bookstore_utils-go/rest_errors"
)

var ItemsService itemsServiceInterface = &itemsService{}

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, resterrs.RestErr)
	Get(string) (*items.Item, resterrs.RestErr)
	Search(queries.EsQuery) ([]items.Item, resterrs.RestErr)
}

type itemsService struct{}

func NewService() itemsServiceInterface {
	return &itemsService{}
}

func (s *itemsService) Create(item items.Item) (*items.Item, resterrs.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil

}

func (s *itemsService) Get(id string) (*items.Item, resterrs.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Search(query queries.EsQuery) ([]items.Item, resterrs.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}
