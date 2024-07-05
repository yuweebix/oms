package pub

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
)

type Producer struct {
	async sarama.AsyncProducer
}

func NewProducer(brokers []string) (*Producer, error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &Producer{async: producer}, nil
}

func (p *Producer) Send(topic string, message any) error {
	msg, err := buildMessage(topic, message)
	if err != nil {
		return err
	}

	fmt.Println("HERE")
	p.async.Input() <- msg

	return nil
}

func (p *Producer) Close() error {
	err := p.async.Close()
	if err != nil {
		return err
	}
	return nil
}

func buildMessage(topic string, message any) (*sarama.ProducerMessage, error) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(jsonMessage),
	}

	return msg, nil
}
