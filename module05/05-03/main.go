package _5_03

import "sync"

type Cache struct {
	storage map[string]int
	mu      sync.RWMutex
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	c.storage[key] += value
	c.mu.Unlock()
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	c.storage[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get(key string) int {
	c.mu.RLock()
	value := c.storage[key]
	c.mu.RUnlock()
	return value
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	delete(c.storage, key)
	c.mu.Unlock()
}
