package pub

import (
	"encoding/json"
	"sync"

	"github.com/IBM/sarama"
)

type Producer struct {
	async sarama.AsyncProducer
	wg    sync.WaitGroup
}

// NewProducer создает нового асинхронного продьюсера Kafka
func NewProducer(brokers []string) (p *Producer, err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true // без вызова Successes будет deadlock
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	p = &Producer{async: producer}

	// проверяем на правильную работу
	go func() {
		for range p.async.Successes() {
			p.wg.Done()
		}
	}()

	go func() {
		for err := range p.async.Errors() {
			panic(err)
		}
	}()

	return
}

// Send отправляет сообщение в заданный топик Kafka.
func (p *Producer) Send(topic string, message any) error {
	msg, err := buildMessage(topic, message)
	if err != nil {
		return err
	}

	p.wg.Add(1)
	p.async.Input() <- msg

	return nil
}

// Close завершает работу продьюсера
func (p *Producer) Close() error {
	p.wg.Wait()
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
