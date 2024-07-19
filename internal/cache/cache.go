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

// extractOrderID хелпер-функция для "вытаскивания" id из ключа
func extractOrderID(orderKey string) (orderID uint64) {
	fmt.Sscanf(orderKey, "order:%d", &orderID)
	return orderID
}
