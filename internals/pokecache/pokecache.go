package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu   *sync.Mutex
	data map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mu:   &sync.Mutex{},
		data: map[string]CacheEntry{},
	}

	go cache.reapLoop(&interval)
	return cache

}

func (c *Cache) reapLoop(interval *time.Duration) {
	ticker := time.NewTicker(*interval)
	for range ticker.C {

		c.mu.Lock()
		for key, entry := range c.data {
			if time.Since(entry.createdAt) >= *interval {
				delete(c.data, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, ok := c.data[key]
	if ok {
		return data.val, ok
	}
	return nil, ok
}
