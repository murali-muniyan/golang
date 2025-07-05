package cache

import (
	"errors"
	"log"
	"time"
)

type Cache[Key comparable, Val any] struct {
	data   map[Key]Val
	maxLen int
}

func New[Key comparable, Val any](maxLen int) *Cache[Key, Val] {
	return &Cache[Key, Val]{
		data:   make(map[Key]Val, 0),
		maxLen: maxLen,
	}
}

func (c *Cache[Key, Val]) AddVal(key Key, val Val) error {
	return c.AddValWithTTL(key, val, -1)
}

func (c *Cache[Key, Val]) AddValWithTTL(key Key, val Val, ttl time.Duration) error {
	if len(c.data) == c.maxLen {
		return errors.New("cache already full")
	}

	if _, exists := c.data[key]; exists {
		return errors.New("data already exists for provided key")
	}

	c.data[key] = val

	if ttl >= 0 {
		go func() {
			time.Sleep(ttl)
			log.Printf("removing data from cache for key: '%v' as ttl expired\n", key)
			c.Remove(key)
		}()
	}

	return nil
}

func (c *Cache[Key, Val]) GetVal(key Key) (val Val, err error) {
	val, exists := c.data[key]
	if !exists {
		return val, errors.New("data not found in cache")
	}

	return val, nil
}

func (c *Cache[Key, Val]) Remove(key Key) (Val, error) {
	val, err := c.GetVal(key)
	if err != nil {
		return val, err
	}

	delete(c.data, key)

	return val, nil
}

func (c *Cache[Key, Val]) IsFull() bool {
	return len(c.data) == c.maxLen
}
