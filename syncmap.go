package yas

import "sync"

// SyncMap is a generic map with string key
type SyncMap[T any] struct {
	m    map[string]T
	sync sync.Mutex
}

// NewSyncMap creates an empty *SyncMap[T]
func NewSyncMap[T any]() *SyncMap[T] {
	return &SyncMap[T]{
		m: make(map[string]T),
	}
}

// Count returns the number of items
func (m *SyncMap[T]) Count() int { return len(m.m) }

// Test uses a Tester to return all passing values' keys
func (m *SyncMap[T]) Test(t Tester[T]) (keys []string) {
	m.sync.Lock()
	for k, v := range m.m {
		if t.Test(v) {
			keys = append(keys, k)
		}
	}
	m.sync.Unlock()
	return
}

// TestFirst uses a Tester to return the first passing value
func (m *SyncMap[T]) TestFirst(t Tester[T]) (key string, val T) {
	m.sync.Lock()
	for k, v := range m.m {
		if t.Test(v) {
			key, val = k, v
			break
		}
	}
	m.sync.Unlock()
	return
}

// Get returns the value for a key
func (m *SyncMap[T]) Get(key string) T { return m.m[key] }

// Remove deletes keys
func (m *SyncMap[T]) Remove(keys ...string) {
	m.sync.Lock()
	for _, key := range keys {
		delete(m.m, key)
	}
	m.sync.Unlock()
}

// RemoveTest deletes keys that pass the test
func (m *SyncMap[T]) RemoveTest(t Tester[T]) {
	keys := make([]string, 0)
	m.sync.Lock()
	for k, v := range m.m {
		if t.Test(v) {
			keys = append(keys, k)
		}
	}
	for _, key := range keys {
		delete(m.m, key)
	}
	m.sync.Unlock()
}

// Set changes the value for a key
func (m *SyncMap[T]) Set(key string, t T) {
	m.sync.Lock()
	m.m[key] = t
	m.sync.Unlock()
}

// Sync calls a function inside the mutex lock state
func (m *SyncMap[T]) Sync(f func()) {
	m.sync.Lock()
	f()
	m.sync.Unlock()
}

// Each calls a function, once for every value, inside the mutex lock state
func (m *SyncMap[T]) Each(f func(id string, v T)) {
	m.sync.Lock()
	for k, v := range m.m {
		f(k, v)
	}
	m.sync.Unlock()
}
