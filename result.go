package box

// Result represents the outcome of an operation that can either succeed with a value of type T
// or fail with an error. It wraps the standard built-in error type.
type Result[T any] interface {
	// IsSuccess returns true if the operation completed successfully, otherwise false.
	IsSuccess() bool

	// IsFailure returns true if the operation failed with an error, otherwise false.
	IsFailure() bool

	// Value returns the success value if present, or the zero value of type T.
	Value() T

	// Error returns the underlying error if it failed, or nil.
	Error() error
}

// Success struct represents a successful outcome within a [Result] container.
type Success[T any] struct {
	value T
}

// NewSuccess creates and returns a new successful [Result] containing the provided value.
func NewSuccess[T any](value T) Result[T] {
	return &Success[T]{value: value}
}

// IsSuccess always returns true for [Success].
func (s *Success[T]) IsSuccess() bool { return true }

// IsFailure always returns false for [Success].
func (s *Success[T]) IsFailure() bool { return false }

// Value returns the inner wrapped value of [Success].
func (s *Success[T]) Value() T { return s.value }

// Error always returns nil for [Success].
func (s *Success[T]) Error() error { return nil }

// Failure struct represents a failed outcome wrapping an error within a [Result] container.
type Failure[T any] struct {
	err error
}

// NewFailure creates and returns a new failed [Result] containing the provided error.
func NewFailure[T any](err error) Result[T] {
	return &Failure[T]{err: err}
}

// IsSuccess always returns false for [Failure].
func (f *Failure[T]) IsSuccess() bool { return false }

// IsFailure always returns true for [Failure].
func (f *Failure[T]) IsFailure() bool { return true }

// Value returns the zero value of type T since this represents a [Failure].
func (f *Failure[T]) Value() T { var zero T; return zero }

// Error returns the underlying wrapped error of [Failure].
func (f *Failure[T]) Error() error { return f.err }
