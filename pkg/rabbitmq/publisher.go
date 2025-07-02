package rabbitmq

import (
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
	"github.com/streadway/amqp"
	"github.com/vmdt/gogameserver/pkg/logger"
)

type IPublisher interface {
	PublishMessage(msg interface{}, name string, key string) error
}

type Publisher struct {
	cfg  *RabbitMQConfig
	conn *amqp.Connection
	log  logger.ILogger
	ctx  context.Context
}

func (p *Publisher) PublishMessage(msg interface{}, name string, key string) error {
	data, err := jsoniter.Marshal(msg)

	if err != nil {
		p.log.Error("Error in marshalling message to publish message")
		return err
	}

	channel, err := p.conn.Channel()
	if err != nil {
		p.log.Error("Error in opening channel to consume message")
		return err
	}

	defer channel.Close()

	err = channel.ExchangeDeclare(
		name,       // name
		p.cfg.Kind, // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)

	publishingMsg := amqp.Publishing{
		Body:         data,
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		MessageId:    uuid.NewV4().String(),
		Timestamp:    time.Now(),
	}

	err = channel.Publish(
		name,  // exchange
		key,   // routing key
		false, // mandatory
		false, // immediate
		publishingMsg,
	)

	if err != nil {
		p.log.Error("Error in publishing message")
		return err
	}

	p.log.Infof("Published message: %s", publishingMsg.Body)
	return nil
}

func NewPublisher(ctx context.Context, cfg *RabbitMQConfig, conn *amqp.Connection, log logger.ILogger) IPublisher {
	return &Publisher{ctx: ctx, cfg: cfg, conn: conn, log: log}
}
