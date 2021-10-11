package elasticsearch

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

var (
	EsClient esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(c *elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
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

func (c *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}
