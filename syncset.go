package yas

import "sync"

// SyncSet is a map[T]struct{} with sync.Mutex
type SyncSet[T comparable] struct {
	m    map[T]struct{}
	sync sync.Mutex
}

// NewSyncSet creates an empty SyncSet[T]
func NewSyncSet[T comparable]() *SyncSet[T] {
	return &SyncSet[T]{
		m: make(map[T]struct{}),
	}
}

// Has checks value is in Set
func (ss *SyncSet[T]) Has(t T) bool {
	_, ok := ss.m[t]
	return ok
}

// Add stores a value
func (s *SyncSet[T]) Add(t T) {
	s.sync.Lock()
	s.m[t] = emptyStruct
	s.sync.Unlock()
}

// Remove deletes a value
func (s *SyncSet[T]) Remove(t T) {
	s.sync.Lock()
	delete(s.m, t)
	s.sync.Unlock()
}

// Slice returns this SyncSet[T] as []T
func (s *SyncSet[T]) Slice() []T {
	s.sync.Lock()
	i, slice := 0, make([]T, len(s.m))
	for v := range s.m {
		slice[i] = v
		i++
	}
	s.sync.Unlock()
	return slice
}
