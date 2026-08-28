package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]cacheEntry
	mutex sync.RWMutex
}

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	csh := Cache{
		Entry: make(map[string]cacheEntry),
		mutex: sync.RWMutex{},
	}
	go csh.reapLoop(interval)
	return &csh
}

func (cash *Cache) Add(key string, Value []byte) {
	cash.mutex.Lock()
	defer cash.mutex.Unlock()

	cash.Entry[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       Value,
	}

}

func (cash *Cache) Get(key string) ([]byte, bool) {
	cash.mutex.RLock()
	defer cash.mutex.RUnlock()
	entry, ok := cash.Entry[key]
	if !ok {
		return nil, false
	}
	value := entry.Val
	return value, true
}

func (cash *Cache) reapLoop(timeout time.Duration) {
	ticker := time.NewTicker(timeout)
	defer ticker.Stop()
	for range ticker.C {
		cash.mutex.Lock()
		for k, v := range cash.Entry {
			timeSince := time.Since(v.CreatedAt)
			if timeSince >= timeout {
				delete(cash.Entry, k)
			}
		}
		cash.mutex.Unlock()
	}
}
