package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           sync.Mutex
	ticker       *time.Ticker
	ttl          time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	content   []byte
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Println("adding to cache")
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		content:   val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Println("getting from cache")
	c.mu.Lock()
	defer c.mu.Unlock()
	val, exists := c.cacheEntries[key]
	if !exists {
		fmt.Println("not found in cache")
		return nil, false
	}
	fmt.Println("found in cache")
	return val.content, true
}

func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		ticker:       time.NewTicker(ttl),
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) reapLoop() {
	for range c.ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.cacheEntries {
			if now.Sub(entry.createdAt) >= c.ttl {
				fmt.Printf("deleting %s from cache\n", key)
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
