package server

import (
	"sync"
	"time"
)

type cacheItem struct {
	value      interface{}
	expireTime time.Time
}

// MemoryCache 内存缓存，支持 TTL 过期
type MemoryCache struct {
	mu    sync.RWMutex
	items map[string]*cacheItem
}

// NewMemoryCache 创建新的内存缓存实例
func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{items: make(map[string]*cacheItem)}
	go c.cleaner()
	return c
}

// Get 获取缓存值
func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	if time.Now().After(item.expireTime) {
		return nil, false
	}
	return item.value, true
}

// Set 设置缓存值
func (c *MemoryCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = &cacheItem{
		value:      value,
		expireTime: time.Now().Add(ttl),
	}
}

// Del 删除缓存
func (c *MemoryCache) Del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Clear 清空所有缓存
func (c *MemoryCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*cacheItem)
}

// 定期清理过期缓存
func (c *MemoryCache) cleaner() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.items {
			if now.After(v.expireTime) {
				delete(c.items, k)
			}
		}
		c.mu.Unlock()
	}
}
