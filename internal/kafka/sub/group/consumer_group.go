package group

import (
	"context"
	"errors"
	"sync"

	"github.com/IBM/sarama"
)

type Group struct {
	client      sarama.ConsumerGroup
	wg          *sync.WaitGroup
	ready       chan bool
	messageChan chan<- string
}

// NewConsumerGroup инициализирует новую Kafka группу консьюмеров
func NewConsumerGroup(brokers, topics []string, groupID string, messageChan chan<- string) (*Group, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	client, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	group := &Group{
		client:      client,
		wg:          new(sync.WaitGroup),
		ready:       make(chan bool),
		messageChan: messageChan,
	}

	return group, nil
}

// Start начинает процесс консьюма сообщений
func (group *Group) Start(ctx context.Context, topics []string) error {
	group.wg.Add(1)
	go func() {
		defer group.wg.Done()

		for {
			if err := group.client.Consume(ctx, topics, group); err != nil {
				// при вызове cancelFunc, возвращается ошибка в `Consume`
				if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
					return
				}

				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}

				panic(err)
			}

			group.ready = make(chan bool)
		}
	}()

	<-group.ready
	return nil
}

// Stop останавливает работу группы
func (group *Group) Stop() error {
	group.wg.Wait()
	if err := group.client.Close(); err != nil {
		return err
	}

	return nil
}

// Setup вызывается само, не трогать
func (group *Group) Setup(sarama.ConsumerGroupSession) error {
	close(group.ready)
	return nil
}

// Cleanup вызывается само, не трогать
func (group *Group) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim вызывается само, не трогать
func (group *Group) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			if group.messageChan != nil {
				group.messageChan <- string(message.Value)
			}
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
