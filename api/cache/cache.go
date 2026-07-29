package cache

import (
	"sync"
	"time"
)

type entry struct {
	value     any
	expiresAt time.Time
}

type Store struct {
	mu    sync.RWMutex
	ttl   time.Duration
	items map[string]entry
}

func New(ttl time.Duration) *Store {
	return &Store{ttl: ttl, items: make(map[string]entry)}
}

func (s *Store) Get(key string) (any, bool) {
	s.mu.RLock()
	item, ok := s.items[key]
	s.mu.RUnlock()
	if !ok || time.Now().After(item.expiresAt) {
		if ok {
			s.Delete(key)
		}
		return nil, false
	}
	return item.value, true
}

func (s *Store) Set(key string, value any) {
	s.mu.Lock()
	s.items[key] = entry{value: value, expiresAt: time.Now().Add(s.ttl)}
	s.mu.Unlock()
}

func (s *Store) Delete(keys ...string) {
	s.mu.Lock()
	for _, key := range keys {
		delete(s.items, key)
	}
	s.mu.Unlock()
}
