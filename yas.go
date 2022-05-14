package yas

var emptyStruct = struct{}{}

// Observer is an interface for observing a generic type
type Observer[T any] interface {
	Observe(id string, new, old T)
}

// ObserverFunc is a func type Observer
type ObserverFunc[T any] func(string, T, T)

// Observe implements Observer by calling the func
func (f ObserverFunc[T]) Observe(id string, new, old T) { f(id, new, old) }

// Tester is an interface for testing a generic type
type Tester[T any] interface {
	Test(T) bool
}

// testerFunc is a func type Tester
type TesterFunc[T any] func(T) bool

// Test implements Tester by calling the func
func (f TesterFunc[T]) Test(t T) bool { return f(t) }

// Zero returns the zero value for generic type
func Zero[T any]() (_ T) { return }
