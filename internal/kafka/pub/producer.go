package pub

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

type Producer struct {
	async sarama.AsyncProducer
}

// NewProducer создает нового асинхронного продьюсера Kafka
func NewProducer(brokers []string) (p *Producer, err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true // "If Return.Successes is true, you MUST read from this channel or the Producer will deadlock."
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Producer{async: producer}, nil
}

// Send отправляет сообщение в заданный топик Kafka.
func (p *Producer) Send(topic string, message any) error {
	msg, err := buildMessage(topic, message)
	if err != nil {
		return err
	}

	p.async.Input() <- msg

	// "It is suggested that you send and read messages together in a single select statement."
	select {
	case <-p.async.Successes(): // дефолтный случай: всё получилось, и мы смогли отослать сообщение
	case err := <-p.async.Errors(): // всё плохо :(
		return err
	}

	return nil
}

// Close завершает работу продьюсера
func (p *Producer) Close() error {
	return p.async.Close()
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
