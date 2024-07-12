package pub

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

type Producer struct {
	sync  sarama.SyncProducer
	topic string
}

// NewProducer создает нового синхронного продьюсера Kafka
func NewProducer(brokers []string, topic string) (p *Producer, err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{sync: producer, topic: topic}, nil
}

// Send отправляет сообщение в заданный топик Kafka.
func (p *Producer) Send(message any) error {
	msg, err := buildMessage(p.topic, message)
	if err != nil {
		return err
	}

	_, _, err = p.sync.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

// Close завершает работу продьюсера
func (p *Producer) Close() error {
	return p.sync.Close()
}

// buildMessage переводит message в ProducerMessage
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
