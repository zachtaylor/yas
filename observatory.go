package yas

import "sync"

// Observatory is a generic observable map
type Observatory[T any] struct {
	keygen  func() string
	m       map[string]T
	sync    sync.Mutex
	observe []Observer[T]
}

// NewObservatory creates an empty *Observatory[T]
func NewObservatory[T any](keygen func() string) *Observatory[T] {
	return &Observatory[T]{
		keygen:  keygen,
		m:       make(map[string]T),
		observe: make([]Observer[T], 0),
	}
}

func (this *Observatory[T]) callback(key string, new T, old T) {
	for _, o := range this.observe {
		o(key, new, old)
	}
}

func (this *Observatory[T]) set(key string, t T) {
	this.callback(key, t, this.m[key])
	this.m[key] = t
}

// Count returns the number of items
func (this *Observatory[T]) Count() int { return len(this.m) }

// Filter uses a Tester to return all passing values' keys
func (this *Observatory[T]) Filter(test Tester[T]) (keys []string) {
	this.sync.Lock()
	for k, v := range this.m {
		if test(v) {
			keys = append(keys, k)
		}
	}
	this.sync.Unlock()
	return
}

// First uses a Tester to return the first passing value
func (this *Observatory[T]) First(test Tester[T]) (key string, t T) {
	this.sync.Lock()
	for k, v := range this.m {
		if test(v) {
			key, t = k, v
			break
		}
	}
	this.sync.Unlock()
	return
}

// Get returns the value for a key
func (this *Observatory[T]) Get(key string) T { return this.m[key] }

// Observe adds an observer
func (this *Observatory[T]) Observe(f Observer[T]) { this.observe = append(this.observe, f) }

func (this *Observatory[T]) remove(key string) {
	this.callback(key, Zero[T](), this.m[key])
	delete(this.m, key)
}

// Remove deletes keys
func (this *Observatory[T]) Remove(keys ...string) {
	this.sync.Lock()
	for _, key := range keys {
		this.remove(key)
	}
	this.sync.Unlock()
}

// RemoveTest deletes keys that pass the test
func (this *Observatory[T]) RemoveTest(test Tester[T]) {
	keys := make([]string, 0)
	this.sync.Lock()
	for k, v := range this.m {
		if test(v) {
			keys = append(keys, k)
		}
	}
	for _, key := range keys {
		this.remove(key)
	}
	this.sync.Unlock()
}

// Set changes the value for a key
func (this *Observatory[T]) Set(key string, t T) {
	this.sync.Lock()
	this.set(key, t)
	this.sync.Unlock()
}

// Store creates a new key for the value
func (this *Observatory[T]) Store(t T) (key string) {
	this.sync.Lock()
	for ok := true; ok; _, ok = this.m[key] {
		key = this.keygen()
	}
	this.set(key, t)
	this.sync.Unlock()
	return
}

// Sync calls a function inside the mutex lock state
func (this *Observatory[T]) Sync(f func()) {
	this.sync.Lock()
	f()
	this.sync.Unlock()
}
