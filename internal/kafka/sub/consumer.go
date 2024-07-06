package sub

import (
	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer    sarama.Consumer
	messageChan chan<- string
}

// NewConsumer инициализирует нового Consumer и начинает потребление сообщений из заданного топика
func NewConsumer(brokers []string, topic string, messageChan chan<- string) (c *Consumer, err error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	c = &Consumer{
		consumer:    consumer,
		messageChan: messageChan,
	}

	err = c.begin(topic)
	if err != nil {
		return nil, err
	}

	return
}

// begin начинает потребление сообщений из всех партиций топика и отправляет их в messageChan
func (c *Consumer) begin(topic string) error {
	partitions, err := c.consumer.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		pc, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}

		go func(pc sarama.PartitionConsumer) {
			defer pc.Close()
			for message := range pc.Messages() {
				if c.messageChan != nil {
					c.messageChan <- string(message.Value)
				}
			}
		}(pc)
	}

	return nil
}

// Close закрывает консьюмера и его партиции
func (c *Consumer) Close() error {
	return c.consumer.Close()
}
