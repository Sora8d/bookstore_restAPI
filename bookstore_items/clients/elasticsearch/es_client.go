package elasticsearch

import (
	"context"
	"time"

	"github.com/Sora8d/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

var (
	EsClient esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
	)
	if err != nil {
		panic(err)
	}
	EsClient.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	response, err := c.client.Index().
		Index(index).BodyJson(doc).Type(docType).Do(ctx)

	if err != nil {
		logger.Error("error when trying to index document in es", err)
	}
	return response, err
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {

		logger.Error("error when getting item with given id", err)
		return nil, err
	}
	return result, err
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()

	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
