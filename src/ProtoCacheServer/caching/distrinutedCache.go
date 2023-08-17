package caching

import (
	"fmt"
	"sync"
)

type DistributedCache interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Remove(key string) error
	// Refresh(key string)
}

type memoryDistributedCache struct {
	lock sync.RWMutex
	data map[string][]byte
}

func NewMemoryDistributedCache() DistributedCache {
	return &memoryDistributedCache{
		data: make(map[string][]byte),
	}
}

func (c *memoryDistributedCache) Get(key string) ([]byte, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	val, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("key (%s) not found", key)
	}

	return val, nil
}

func (c *memoryDistributedCache) Set(key string, value []byte) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = value

	// if ttl > 0 {
	// 	go func() {
	// 		<-time.After(ttl)
	// 		delete(c.data, key)
	// 	}()
	// }

	return nil
}

func (c *memoryDistributedCache) Remove(key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.data, key)

	return nil
}
