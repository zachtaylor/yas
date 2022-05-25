package yas

// Stack is a slice with convenience methods
type Stack[T any] []T

// NewStack returns a generic stack pointer
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Pop removes a value from the end of the stack
func (s *Stack[T]) Pop() (t T) {
	if s != nil {
		if slice := *s; len(slice) > 0 {
			t, *s = slice[0], slice[1:]
		}
	}
	return
}

// Push adds values to the end of the stack
func (s *Stack[T]) Push(t ...T) {
	if s != nil {
		*s = append(*s, t...)
	}
}

// Shift removes a value from the front of the stack
func (s *Stack[T]) Shift() (t T) {
	if s != nil {
		if slice := *s; len(slice) > 0 {
			t, *s = Shift(slice)
		}
	}
	return
}

// Unshift adds values to the front of the stack
func (s *Stack[T]) Unshift(t ...T) {
	if s != nil {
		*s = append(*s, t...)
	}
}

// Each calls a function once for every value
func (s *Stack[T]) Each(f func(int, T)) {
	if s != nil {
		for i, v := range *s {
			f(i, v)
		}
	}
}
