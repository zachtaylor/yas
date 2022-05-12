package yas

import "sync"

// Observatory is a generic observable map
type Observatory[T comparable] struct {
	keygen  func() string
	m       map[string]T
	sync    sync.Mutex
	observe []Observer[T]
}

// NewObservatory creates an empty *Observatory[T]
func NewObservatory[T comparable](keygen func() string) *Observatory[T] {
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
	if Zero[T]() == t {
		delete(this.m, key)
	} else {
		this.m[key] = t
	}
}

// Count returns the number of items
func (this *Observatory[T]) Count() int { return len(this.m) }

// Filter uses a Tester to return all passing values
func (this *Observatory[T]) Filter(test Tester[T]) (ts []T) {
	this.sync.Lock()
	for _, v := range this.m {
		if test(v) {
			ts = append(ts, v)
			break
		}
	}
	this.sync.Unlock()
	return
}

// First uses a Tester to return the first passing value
func (this *Observatory[T]) First(test Tester[T]) (t T) {
	this.sync.Lock()
	for _, v := range this.m {
		if test(v) {
			t = v
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

// Remove deletes a key
func (this *Observatory[T]) Remove(key string) { this.Set(key, Zero[T]()) }

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
