package cache

import "time"

type CacheI[Key comparable, Val any] interface {
	AddVal(key Key, val Val) error
	AddValWithTTL(key Key, val Val, ttl time.Duration) error
	GetVal(key Key) (Val, error)
	Remove(key Key) (Val, error)
	IsFull() bool
}
