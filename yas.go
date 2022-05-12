package yas

// Tester is a func type test of a generic type
type Tester[T any] func(T) bool

// Observer is a func type change handler
type Observer[T any] func(id string, new, old T)

// Zero returns the zero value for generic type
func Zero[T any]() (_ T) { return }
