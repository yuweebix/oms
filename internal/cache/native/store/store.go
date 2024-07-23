package store

import (
	"sync"
	"time"
)

type entry struct {
	value     any
	duration  *time.Duration
	createdAt time.Time
}

// Store хранилище данных вида пар ключ-значение
type Store struct {
	mu   *sync.RWMutex
	data map[string]entry
	stop chan struct{}
}

// NewStore конструктор, что задаёт хранилище, а также начинает очищающую горутину
func NewStore(cleanupTickrate time.Duration) *Store {
	s := &Store{
		mu:   &sync.RWMutex{},
		data: make(map[string]entry),
		stop: make(chan struct{}),
	}
	go s.cleanup(cleanupTickrate)
	return s
}

// Set задаёт значение с опциональным TTL
func (s *Store) Set(key string, value any, ttl *time.Duration) {
	s.mu.Lock()
	s.data[key] = entry{value: value, duration: ttl, createdAt: time.Now()}
	s.mu.Unlock()
}

// Get возващает непросроченный элемент
func (s *Store) Get(key string) (value any, ok bool) {
	s.mu.RLock()
	e, exists := s.data[key]
	s.mu.RUnlock()
	if !exists {
		return nil, false
	}
	if e.duration != nil && time.Since(e.createdAt) > *e.duration {
		s.Del(key)
		return nil, false
	}
	return e.value, true
}

// Del удаляет по ключу
func (s *Store) Del(key string) {
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()
}

// cleanup занимается очисткой хранилища от просроченных записей
func (s *Store) cleanup(cleanupTickrate time.Duration) {
	ticker := time.NewTicker(cleanupTickrate)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			for k, e := range s.data {
				if e.duration == nil {
					continue
				}
				if time.Since(e.createdAt) > *e.duration {
					delete(s.data, k)
				}
			}
			s.mu.Unlock()
		case <-s.stop:
			return
		}
	}
}

// StopCleanup останавливает процесс очистки хранилища от просроченных записей
func (s *Store) StopCleanup() {
	s.stop <- struct{}{}
	close(s.stop)
}
