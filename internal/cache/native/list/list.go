package list

import (
	"sync"
	"sync/atomic"
)

// List потокобезопасный список
type List[T comparable] struct {
	mu   *sync.RWMutex
	len  atomic.Int32
	data map[T]struct{}
}

// NewList конструктор
func NewList[T comparable]() *List[T] {
	return &List[T]{
		mu:   &sync.RWMutex{},
		data: make(map[T]struct{}),
	}
}

// Add добавляет элемент в список
func (l *List[T]) Add(el T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, ok := l.data[el]
	if !ok {
		l.len.Add(1)
	}
	l.data[el] = struct{}{}
}

// Rem удаляет элемент из списка
func (l *List[T]) Rem(el T) {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, ok := l.data[el]
	if ok {
		l.len.Add(-1)
	}
	delete(l.data, el)
}

// Get возвращает копию полного списка в форме мапы
func (l *List[T]) Get() map[T]struct{} {
	l.mu.RLock()
	defer l.mu.RUnlock()
	copy := make(map[T]struct{})
	for k, v := range l.data {
		copy[k] = v
	}
	return copy
}

// Len длина списка
func (l *List[T]) Len() int32 {
	return l.len.Load()
}
