package box

import (
	"sync/atomic"
)

// Atomic is a type-safe atomic value container.
// The zero value is the zero value of T.
// An Atomic must not be copied after first use.
type Atomic[T any] struct {
	_    noCopy
	zero T
	v    atomic.Value
}

// Load atomically loads and returns the value stored in x.
// If the stored value is not of type T, returns the zero value of T.
func (s *Atomic[T]) Load() T {
	if val, ok := s.v.Load().(T); ok {
		return val
	}
	return s.zero
}

// Store atomically stores val into x.
func (s *Atomic[T]) Store(val T) {
	s.v.Store(val)
}

// Swap atomically stores new into x and returns the previous value.
// If the previous value is not of type T, returns the zero value of T.
func (s *Atomic[T]) Swap(new T) (old T) {
	if old, ok := s.v.Swap(new).(T); ok {
		return old
	}
	return s.zero
}

// CompareAndSwap executes the compare-and-swap operation for x.
// It returns true if the swap was successful (old matched the current value),
// false otherwise.
func (s *Atomic[T]) CompareAndSwap(old, new T) (swapped bool) {
	return s.v.CompareAndSwap(old, new)
}
