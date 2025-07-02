package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	ExchangeName string `mapstructure:"exchange_name"`
	Kind         string `mapstructure:"kind"`
	Uri          string `mapstructure:"uri"`
}

func NewRabbitMQConn(cfg *RabbitMQConfig, ctx context.Context) (*amqp.Connection, error) {
	var connAddr string
	if cfg.Uri == "" {
		connAddr = fmt.Sprintf(
			"amqp://%s:%s@%s:%d/",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
		)
	} else {
		connAddr = cfg.Uri
	}

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 10 * time.Second // Maximum time to retry
	maxRetries := 5                      // Number of retries (including the initial attempt)

	var conn *amqp.Connection
	var err error

	err = backoff.Retry(func() error {

		conn, err = amqp.Dial(connAddr)
		if err != nil {
			log.Errorf("Failed to connect to RabbitMQ: %v. Connection information: %s", err, connAddr)
			return err
		}

		return nil
	}, backoff.WithMaxRetries(bo, uint64(maxRetries-1)))

	log.Info("Connected to RabbitMQ")

	go func() {
		select {
		case <-ctx.Done():
			err := conn.Close()
			if err != nil {
				log.Error("Failed to close RabbitMQ connection")
			}
			log.Info("RabbitMQ connection is closed")
		}
	}()

	return conn, err
}
