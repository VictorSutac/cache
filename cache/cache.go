package cache

import (
	"sync"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()         // Блокировка для записи
	defer c.mu.Unlock() // Освобождение блокировки после записи
	c.data[key] = value
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()         // Блокировка для чтения
	defer c.mu.RUnlock() // Освобождение блокировки после чтения
	return c.data[key]
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()         // Блокировка для удаления
	defer c.mu.Unlock() // Освобождение блокировки после удаления
	delete(c.data, key)
}
