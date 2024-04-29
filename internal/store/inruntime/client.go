package inruntime

import (
	"sync"
)

type Store interface {
	Get(key string) (string, bool)
	Set(key, val string)
	IsExpire(key, val string)
}

// TODO - use `redis` or clear map by `copy`! (without `copy` GC will not free memory!); `map` has this issue.
var urls = make(map[string]string)

// StoreMem - store URLs in memory
type StoreMem struct {
	sync.Mutex
}

func New() *StoreMem {
	return &StoreMem{}
}

// Get - get value by key.
func (s *StoreMem) Get(key string) (string, bool) {
	s.Lock()
	defer s.Unlock()

	val, ok := urls[key]

	if s.Expire(key) {
		return val, false
	}

	return val, ok
}

// Set - set value by key.
func (s *StoreMem) Set(key, val string) {
	s.Lock()
	defer s.Unlock()

	urls[key] = val
}

func (s *StoreMem) Expire(key string) bool {
	// TODO it is strange use url-shortener without expiration. Expiration may be 1 or 10 years, but it must exist.
	return false
}
