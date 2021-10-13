package items

import (
	"bookstoreapi/items/clients/elasticsearch"
	"bookstoreapi/items/domain/queries"
	"encoding/json"
	"errors"
	"strings"

	"github.com/Sora8d/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.EsClient.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.EsClient.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError("no item found with given id")
		}
		return rest_errors.NewInternalServerError("error when trying to get item with given id", errors.New("database error"))
	}
	bytes, _ := result.Source.MarshalJSON()
	if err := json.Unmarshal(bytes, &i); err != nil {
		return rest_errors.NewInternalServerError("error parsing db response", errors.New("database error"))
	}

	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, rest_errors.RestErr) {
	result, err := elasticsearch.EsClient.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to serch documents", errors.New("db error"))
	}
	items := make([]Item, result.TotalHits())

	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerError("error trying to parse response", errors.New("database error"))
		}
		item.Id = hit.Id
		items[index] = item
	}
	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundError("No items with that criteria")
	}
	return items, nil
}
