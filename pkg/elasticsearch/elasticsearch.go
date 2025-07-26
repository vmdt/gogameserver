package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type ElasticOptions struct {
	URL string `mapstructure:"url"`
}

func NewElasticClient(options *ElasticOptions) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{options.URL},
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
