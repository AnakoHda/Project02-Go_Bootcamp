package main

import (
	"container/list"
	"errors"
	"sync"
)

type dataKey[K comparable, T any] struct {
	data T
	key  K
}
type Cache[K comparable, T any] struct {
	mutex    sync.RWMutex
	list     *list.List
	items    map[K]*list.Element
	capacity int
}

// храним лист на элементы которого указывает map
// через ключ K получаем доступ к конкретному элементу за O(1)
func NewCache[K comparable, T any](capacity int) (*Cache[K, T], error) {
	if capacity <= 0 {
		return nil, errors.New("capacity too small")
	}
	return &Cache[K, T]{
		capacity: capacity,
		list:     list.New(),
		items:    make(map[K]*list.Element, capacity),
	}, nil
}
func (c *Cache[K, T]) Set(key K, value T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.items[key]; ok {
		element.Value.(*dataKey[K, T]).data = value
		c.list.MoveToFront(element)

	} else if c.list.Len() >= c.capacity {
		delete(c.items, c.list.Back().Value.(*dataKey[K, T]).key)
		c.list.Remove(c.list.Back())

		c.list.PushFront(&dataKey[K, T]{key: key, data: value})
		c.items[key] = c.list.Front()
	} else {
		c.list.PushFront(&dataKey[K, T]{key: key, data: value})
		c.items[key] = c.list.Front()
	}
}
func (c *Cache[K, T]) Get(key K) (T, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element, ok := c.items[key]; ok {
		c.list.MoveToFront(element)
		return element.Value.(*dataKey[K, T]).data, true
	}

	var zero T
	return zero, false
}
func (c *Cache[K, T]) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.list.Init()
	c.items = make(map[K]*list.Element, c.capacity)
}
