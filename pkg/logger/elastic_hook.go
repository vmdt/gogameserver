package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
)

type ElasticHook struct {
	client    *elasticsearch.Client
	indexName string
}

func NewElasticHook(client *elasticsearch.Client, indexName string) *ElasticHook {
	return &ElasticHook{
		client:    client,
		indexName: indexName,
	}
}

func (h *ElasticHook) Fire(entry *logrus.Entry) error {
	logData := make(map[string]interface{})

	// Add all fields (entry.Data)
	for k, v := range entry.Data {
		logData[k] = v
	}

	logData["message"] = entry.Message
	logData["level"] = entry.Level.String()
	logData["timestamp"] = entry.Time.Format(time.RFC3339)

	// Marshal log to JSON
	body, err := json.Marshal(logData)
	if err != nil {
		return err
	}

	// Push to Elasticsearch
	_, err = h.client.Index(h.indexName, bytes.NewReader(body), h.client.Index.WithContext(context.Background()))
	return err
}

func (h *ElasticHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
