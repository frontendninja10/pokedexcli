package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mu *sync.Mutex
	interval time.Duration
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.entry[key]
	if exists {
		return entry.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for t := range ticker.C {
        c.mu.Lock()
        for k, v := range c.entry {
            if t.Sub(v.createdAt) > c.interval {
                delete(c.entry, k)
            }
        }
        c.mu.Unlock()
    }
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entry: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
		interval: interval,
		
	}

	go c.reapLoop()

	return c
}