package sub

import (
	"errors"
	"strings"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer   sarama.Consumer
	partitions []sarama.PartitionConsumer
}

// NewConsumer инициализирует нового Consumer
func NewConsumer(brokers []string) (c *Consumer, err error) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{consumer: consumer}, nil
}

// Start начинает потребление сообщений из всех партиций топика и отправляет их в messageChan
func (c *Consumer) Start(topic string, messageChan chan<- string) error {
	partitions, err := c.consumer.Partitions(topic)
	if err != nil {
		return err
	}

	for _, partition := range partitions {
		pc, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}

		c.partitions = append(c.partitions, pc) // потом нужно будет для остановки

		// Messages канал закроется при остановке -> горутина завершится
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				if messageChan != nil {
					messageChan <- string(msg.Value)
				}
			}
		}(pc)
	}

	return nil
}

// Stop закрывает все партиции и возвращает общую ошибку, если какие-то партиции не удалось закрыть
func (c *Consumer) Stop() error {
	var errMsgs []string

	// по-партиционно закрывает и собираем ошибки
	for _, pc := range c.partitions {
		if err := pc.Close(); err != nil {
			errMsgs = append(errMsgs, err.Error())
		}
	}

	if len(errMsgs) > 0 {
		return errors.New(strings.Join(errMsgs, "; "))
	}

	return nil
}

// Close закрывает консьюмера
func (c *Consumer) Close() error {
	return c.consumer.Close()
}
