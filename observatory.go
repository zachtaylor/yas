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

// Observer is a func type change handler
type Observer[T any] func(id string, new, old T)

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

// Get returns the value for a string
func (this *Observatory[T]) Get(key string) T { return this.m[key] }

// Set changes the value for a key
func (this *Observatory[T]) Set(key string, t T) {
	this.sync.Lock()
	this.set(key, t)
	this.sync.Unlock()
}

// Store creates a new key
func (this *Observatory[T]) Store(t T) (key string) {
	this.sync.Lock()
	for ok := true; ok; _, ok = this.m[key] {
		key = this.keygen()
	}
	this.set(key, t)
	this.sync.Unlock()
	return
}

// Remove deletes a string,*T
func (this *Observatory[T]) Remove(k string) { this.Set(k, Zero[T]()) }

// Observe adds an observer
func (this *Observatory[T]) Observe(f Observer[T]) { this.observe = append(this.observe, f) }
