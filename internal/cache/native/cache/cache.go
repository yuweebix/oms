package cache

import (
	"sync"
	"time"

	"gitlab.ozon.dev/yuweebix/homework-1/internal/cache/native/list"
	"gitlab.ozon.dev/yuweebix/homework-1/internal/cache/native/store"
)

// Cache in-memory кеш
type Cache struct {
	mu    *sync.RWMutex
	store store.Store
}

// NewCache конструктор
func NewCache(cleanupTickrate time.Duration) *Cache {
	return &Cache{
		mu:    &sync.RWMutex{},
		store: *store.NewStore(cleanupTickrate),
	}
}

// Close gracefully закрывает хранилище
func (c *Cache) Close() {
	c.store.StopCleanup()
}

// getOrdersList возвращает значение типа список по заданному ключу
func (c *Cache) getOrdersList(key string) *list.List[uint64] {
	listIDs := &list.List[uint64]{}

	// проверяем имеется уже ли лист, и если нет создаём новый
	val, ok := c.store.Get(key)
	switch ok {
	case false:
		listIDs = list.NewList[uint64]()
		c.store.Set(key, listIDs, nil)
	case true:
		listIDs = val.(*list.List[uint64])
	}

	return listIDs
}
