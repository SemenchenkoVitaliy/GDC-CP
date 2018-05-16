package cache

import (
	"time"
)

// cacheItem contains cace value and date it was accessed last time
type cacheItem struct {
	value interface{}
	date  time.Time
}

// cache
type Cache struct {
	cache   map[string]cacheItem
	size    int
	maxSize int
	maxAge  int
}

// NewCache accepts maximal cache size and max age(in seconds) of item in cache
//
// It returns new instance of Cache binded on inut parameters
func NewCache(maxSize int, maxAge int) (c *Cache) {
	return &Cache{
		cache:   make(map[string]cacheItem),
		maxSize: maxSize,
		maxAge:  maxAge,
	}
}

// Add adds new item to cache based on key
func (c *Cache) Add(key string, value interface{}) {
	if c.size > c.maxSize {
		for k, v := range c.cache {
			if time.Now().Sub(v.date) > time.Duration(c.maxAge)*time.Second {
				delete(c.cache, k)
			}
		}
	}
	c.cache[key] = cacheItem{
		value: value,
		date:  time.Now(),
	}
	c.size++
}

// Delete deletes item which corresponds to key in cache
func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}

// Get returns item which corresponds to given key
func (c *Cache) Get(key string) (value interface{}, ok bool) {
	item, ok := c.cache[key]
	c.cache[key] = cacheItem{
		value: value,
		date:  time.Now(),
	}
	return item.value, ok
}
