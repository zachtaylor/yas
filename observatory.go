package yas

import "sync"

// Observatory is a generic observable map
type Observatory[T any] struct {
	m       map[string]T
	sync    sync.Mutex
	observe []Observer[T]
}

// NewObservatory creates an empty *Observatory[T]
func NewObservatory[T any]() *Observatory[T] {
	return &Observatory[T]{
		m:       make(map[string]T),
		observe: make([]Observer[T], 0),
	}
}

func (o *Observatory[T]) callback(key string, new T, old T) {
	for _, o := range o.observe {
		o.Observe(key, new, old)
	}
}

func (o *Observatory[T]) set(key string, val T) {
	o.callback(key, val, o.m[key])
	o.m[key] = val
}

// Count returns the number of items
func (o *Observatory[T]) Count() int { return len(o.m) }

// Test uses a Tester to return all passing values' keys
func (o *Observatory[T]) Test(t Tester[T]) (keys []string) {
	o.sync.Lock()
	for k, v := range o.m {
		if t.Test(v) {
			keys = append(keys, k)
		}
	}
	o.sync.Unlock()
	return
}

// TestFirst uses a Tester to return the first passing value
func (o *Observatory[T]) TestFirst(t Tester[T]) (key string, val T) {
	o.sync.Lock()
	for k, v := range o.m {
		if t.Test(v) {
			key, val = k, v
			break
		}
	}
	o.sync.Unlock()
	return
}

// Get returns the value for a key
func (o *Observatory[T]) Get(key string) T { return o.m[key] }

// Observe adds an observer
func (o *Observatory[T]) Observe(f Observer[T]) { o.observe = append(o.observe, f) }

func (o *Observatory[T]) remove(key string) {
	o.callback(key, o.m[key], Zero[T]())
	delete(o.m, key)
}

// Remove deletes keys
func (o *Observatory[T]) Remove(keys ...string) {
	o.sync.Lock()
	for _, key := range keys {
		o.remove(key)
	}
	o.sync.Unlock()
}

// RemoveTest deletes keys that pass the test
func (o *Observatory[T]) RemoveTest(t Tester[T]) {
	keys := make([]string, 0)
	o.sync.Lock()
	for k, v := range o.m {
		if t.Test(v) {
			keys = append(keys, k)
		}
	}
	for _, key := range keys {
		o.remove(key)
	}
	o.sync.Unlock()
}

// Set changes the value for a key
func (o *Observatory[T]) Set(key string, t T) {
	o.sync.Lock()
	o.set(key, t)
	o.sync.Unlock()
}

// Sync calls a function inside the mutex lock state
func (o *Observatory[T]) Sync(f func()) {
	o.sync.Lock()
	f()
	o.sync.Unlock()
}
