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

func NewProducer(brokers []string) (p *Producer, err error) {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	p = &Producer{async: producer}

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

func (p *Producer) Send(topic string, message any) error {
	msg, err := buildMessage(topic, message)
	if err != nil {
		return err
	}

	p.wg.Add(1)
	p.async.Input() <- msg

	return nil
}

func (p *Producer) Close() error {
	p.wg.Wait()
	return p.async.Close()
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
