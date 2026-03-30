package cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	data      []byte
	fetchedAt time.Time
}

type Cache struct {
	lock    sync.RWMutex
	entries map[string]CacheEntry
	ttl     time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		entries: make(map[string]CacheEntry),
		ttl:     ttl,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	// Create lock for reading cache
	cache.lock.RLock()
	defer cache.lock.RUnlock()

	// Getting cache entry
	entry, ok := cache.entries[key]
	if !ok || time.Since(entry.fetchedAt) > cache.ttl {
		return nil, false
	}
	return entry.data, true
}

func (cache *Cache) Set(key string, data []byte) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	cache.entries[key] = CacheEntry{
		data:      data,
		fetchedAt: time.Now(),
	}
}
