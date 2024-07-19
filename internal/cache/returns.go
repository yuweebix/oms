package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/models"
)

// GetReturns позволяет получить возвраты из кеша
func (c *Cache) GetReturns(ctx context.Context, limit uint64, offset uint64) (list []*models.Order, err error) {
	statusKey := fmt.Sprintf("orders:by_status:%s", models.StatusReturned)

	// limit можно использовать только при BYSCORE или BYLEX, так что в простом случае просто предпосчитать
	start := offset
	stop := offset + limit - 1

	orderKeys, err := c.client.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:   statusKey,
		Start: start,
		Stop:  stop,
	}).Result()
	if err != nil {
		return nil, err
	}

	for _, key := range orderKeys {
		// нужен лишь id заказа для получения всего заказа
		orderID := extractOrderID(key)
		order, err := c.GetOrder(ctx, &models.Order{ID: orderID})
		if err != nil {
			return nil, err
		}

		if order == nil {
			continue
		}

		list = append(list, order)
	}

	return list, nil
}
