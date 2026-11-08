package pokecache

import "sync"

type Cache struct {
	mu   *sync.RWMutex
	data map[string][]byte
}

func NewCache() Cache {
	return Cache{
		mu:   &sync.RWMutex{},
		data: map[string][]byte{},
	}
}

func (c Cache) Add(key string, val []byte) {
	c.data[key] = val
}

func (c Cache) Get(key string) ([]byte, bool) {
	data, ok := c.data[key]
	if ok {
		return data, ok
	}
	return nil, ok
}
