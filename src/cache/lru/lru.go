package lru

import "container/list"

type lruCache struct {
	size       int
	cache_list *list.List
	cache_map  map[string]*list.Element
}

type Entry struct {
	key   string
	value interface{}
}

func NewLruCache(size int) *lruCache {
	return &lruCache{
		size:       size,
		cache_list: list.New(),
		cache_map:  make(map[string]*list.Element),
	}
}

func (cache *lruCache) Len() int {
	return cache.cache_list.Len()
}

func (c *lruCache) Add(key string, value interface{}) {
	if e, ok := c.cache_map[key]; ok {
		c.cache_list.MoveToFront(e)
		e.Value.(*Entry).value = value
		return
	}
	entry := &Entry{key, value}
	e := c.cache_list.PushFront(entry)
	c.cache_map[key] = e
	// if cache size is full
	for c.Len() > c.size {
		e := c.cache_list.Back()
		if e != nil {
			c.cache_list.Remove(e)
			delete(c.cache_map, e.Value.(*Entry).key)
		}
	}
}

func (c *lruCache) Get(key string) interface{} {
	if e, ok := c.cache_map[key]; ok {
		c.cache_list.MoveToFront(e)
		return e.Value.(*Entry).value
	}
	return nil
}

func (c *lruCache) Remove(key string) {
	if e, ok := c.cache_map[key]; ok {
		c.cache_list.Remove(e)
		delete(c.cache_map, e.Value.(*Entry).key)
	}
}
