package yas

// Set is a map[T]struct{}
type Set[T comparable] map[T]struct{}

var emptyStruct = struct{}{}

// Has checks value is in Set
func (s Set[T]) Has(t T) bool {
	_, ok := s[t]
	return ok
}

// Add stores a value
func (s Set[T]) Add(t T) { s[t] = emptyStruct }

// Remove deletes a value
func (s Set[T]) Remove(t T) { delete(s, t) }

// Slice returns this Set[T] as []T
func (s Set[T]) Slice() []T {
	i, slice := 0, make([]T, len(s))
	for v := range s {
		slice[i] = v
		i++
	}
	return slice
}
