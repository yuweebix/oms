package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Cache представляет собой структуру redis-клиента
type Cache struct {
	client *redis.Client
}

// NewCache конструктор Cache
func NewCache(addr string, password string, db int) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &Cache{client}, nil
}

// Close для закрытия подключения
func (c *Cache) Close() error {
	return c.client.Close()
}

// FlushAll очищает бд
func (c *Cache) FlushAll(ctx context.Context) error {
	return c.client.FlushAll(ctx).Err()
}

// zRemOrdersByStatus удаляет заказ из orders:by_status:* упорядоченных множеств
func (c *Cache) zRemOrdersByStatus(ctx context.Context, orderKey string) error {
	status, err := c.client.HGet(ctx, orderKey, "status").Result()
	switch err {
	// ничего не вернули -> пытаемся удалить отовсюду
	case redis.Nil:
		// сначала получаем множества
		keys, err := c.client.Keys(ctx, "orders:by_status:*").Result()
		if err != nil {
			return err
		}
		// затем удаляем из каждого
		for _, k := range keys {
			err := c.client.ZRem(ctx, k, orderKey).Err()
			if err != nil {
				return err
			}
		}
	// статус вернулся -> надо удалить только тот, что передался
	case nil:
		err = c.client.ZRem(ctx, fmt.Sprintf("orders:by_status:%s", status), orderKey).Err()
		if err != nil {
			return err
		}
	// какая-то другая внутренняя ошибка
	default:
		return err
	}

	return nil
}

// zRemUserOrders удяляет заказ из `user:*:orders“ упорядоченных множеств
func (c *Cache) zRemUserOrders(ctx context.Context, orderKey string) error {
	// т.к. вызывается только при TTL, то HGET никогда не вернет юзера

	// сначала получаем множества
	keys, err := c.client.Keys(ctx, "user:*:orders").Result()
	if err != nil {
		return err
	}

	// затем удаляем из каждого
	for _, k := range keys {
		err := c.client.ZRem(ctx, k, orderKey).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

// extractOrderID хелпер-функция для "вытаскивания" id из ключа
func extractOrderID(orderKey string) (orderID uint64) {
	fmt.Sscanf(orderKey, "order:%d", &orderID)
	return orderID
}
