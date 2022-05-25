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

// Reverse changes the index of all values in-place
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Shift removes the first element of a slice
func Shift[T any](slice []T) (T, []T) { return slice[0], slice[1:] }

// Tester is an interface for testing a generic type
type Tester[T any] interface {
	Test(T) bool
}

// testerFunc is a func type Tester
type TesterFunc[T any] func(T) bool

// Test implements Tester by calling the func
func (f TesterFunc[T]) Test(t T) bool { return f(t) }

// Unshift adds elements to the front of a slice
func Unshift[T any](slice []T, t ...T) []T { return append(t, slice...) }

// Zero returns the zero value for generic type
func Zero[T any]() (_ T) { return }
