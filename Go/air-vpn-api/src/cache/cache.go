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
	Entries map[string]CacheEntry
	Ttl     time.Duration
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	// Create lock for reading cache
	cache.lock.RLock()
	defer cache.lock.RUnlock()

	// Getting cache entry
	entry, ok := cache.Entries[key]
	if !ok || time.Since(entry.fetchedAt) > cache.Ttl {
		return nil, false
	}
	return entry.data, true
}

func (cache *Cache) Set(key string, data []byte) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	cache.Entries[key] = CacheEntry{
		data:      data,
		fetchedAt: time.Now(),
	}
}
