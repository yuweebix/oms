package sub

import (
	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer    sarama.Consumer
	messageChan chan<- string
	partitions  []sarama.PartitionConsumer
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

		c.partitions = append(c.partitions, pc) // будут нужны для закрытия

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				c.messageChan <- string(message.Value)
			}
		}(pc)
	}

	return nil
}

// Close закрывает консьюмера и его партиции
func (c *Consumer) Close() error {
	for _, pc := range c.partitions {
		if err := pc.Close(); err != nil {
			return err
		}
	}
	return c.consumer.Close()
}
