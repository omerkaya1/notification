package mq

import (
	"fmt"
	"github.com/omerkaya1/notification/internal/config"
	"github.com/omerkaya1/notification/internal/errors"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
)

// MessageQueue .
type MessageQueue struct {
	Conn *amqp.Connection
	conf *config.Config
}

// NewMessageQueue .
func NewMessageQueue(conf *config.Config) (*MessageQueue, error) {
	if conf == nil || conf.Host == "" || conf.Port == "" || conf.User == "" || conf.Password == "" {
		return nil, errors.ErrBadQueueConfiguration
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port))
	if err != nil {
		return nil, err
	}
	return &MessageQueue{Conn: conn, conf: conf}, nil
}

// EmmitMessages .
func (mq *MessageQueue) EmmitMessages() error {
	ch, err := mq.Conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		"events",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// Handle interrupt
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt)

MQ:
	for {
		select {
		case <-exitChan:
			log.Println("Exit the programme.")
			ch.Close()
			mq.Conn.Close()
			break MQ
		case d, ok := <-msgs:
			if !ok {
				break MQ
			}
			log.Println(string(d.Body))
		}
	}
	return nil
}
