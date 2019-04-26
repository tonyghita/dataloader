package dataloader

import "sync"

// InMemoryCache is an in memory implementation of Cache interface.
// this simple implementation is well suited for
// a "per-request" dataloader (i.e. one that only lives
// for the life of an http request) but it not well suited
// for long lived cached items.
type InMemoryCacheSyncMap struct {
	mu    sync.Mutex
	items *sync.Map
}

// NewCache constructs a new InMemoryCache
func NewSyncMapCache() *InMemoryCacheSyncMap {
	return &InMemoryCacheSyncMap{
		items: &sync.Map{},
	}
}

// Set sets the `value` at `key` in the cache
func (c *InMemoryCacheSyncMap) Set(key string, value Thunk) {
	c.items.Store(key, value)
}

// Get gets the value at `key` if it exsits, returns value (or nil) and bool
// indicating of value was found
func (c *InMemoryCacheSyncMap) Get(key string) (Thunk, bool) {
	item, found := c.items.Load(key)
	if !found {
		return nil, false
	}

	return item.(Thunk), true
}

// Delete deletes item at `key` from cache
func (c *InMemoryCacheSyncMap) Delete(key string) bool {
	if _, found := c.items.Load(key); found {
		c.items.Delete(key)
		return true
	}
	return false
}

// Clear clears the entire cache
func (c *InMemoryCacheSyncMap) Clear() {
	c.mu.Lock()
	c.items = &sync.Map{}
	c.mu.Unlock()
}
