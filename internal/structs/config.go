package structs

import (
	"sync"
	"time"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

type Config struct {
	Commands map[string]CliCommand
	Cache    *Cache
	PrevUrl  *string
	NextUrl  *string
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	Entries  map[string]CacheEntry
	Duration time.Duration
	Mu       sync.Mutex
}

func NewCache(d time.Duration) *Cache {
	cache := &Cache{Entries: make(map[string]CacheEntry), Duration: d}
	go cache.reaploop(cache.Duration)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entries[key] = CacheEntry{CreatedAt: time.Now(), Val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	cacheEntry, ok := c.Entries[key]
	if ok {
		return cacheEntry.Val, true
	}
	return nil, false
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(c.Duration)
	defer ticker.Stop()

	for range ticker.C {
		for key, entry := range c.Entries {
			c.Mu.Lock()
			if time.Since(entry.CreatedAt) > interval {
				delete(c.Entries, key)
			}
			c.Mu.Unlock()
		}
	}
}
