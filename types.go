package box

// Void represents an empty structure used to signal the absence of data.
// It is commonly used for set implementations or synchronization signals.
type Void = struct{}

// NULL is a globally shared instance of [Void].
var NULL Void

// Optional represents a container object which may or may not contain a non-nil value.
// It provides a type-safe alternative to using nil pointers.
type Optional[T any] interface {
	// IsPresent returns true if there is a value present, otherwise false.
	IsPresent() bool

	// IsEmpty returns true if there is no value present, otherwise false.
	IsEmpty() bool

	// Get returns the value if present. If no value is present, it returns the zero value of type T.
	Get() T
}

// Some represents an [Optional] instance that contains a value.
type Some[T any] struct {
	value T
}

// NewSome creates and returns a new [Optional] containing the provided value.
// The returned instance is a pointer to [Some].
func NewSome[T any](value T) Optional[T] {
	return &Some[T]{value: value}
}

// IsPresent returns true if the receiver is not nil.
func (s *Some[T]) IsPresent() bool { return s != nil }

// IsEmpty returns true if the receiver is nil.
func (s *Some[T]) IsEmpty() bool { return s == nil }

// Get returns the underlying value. If the receiver is nil, it returns the zero value of type T.
func (s *Some[T]) Get() T {
	if s == nil {
		var zero T
		return zero
	}
	return s.value
}

// None represents an [Optional] instance that does not contain any value.
type None[T any] struct{}

// NewNone creates and returns a new [Optional] representing an empty value.
// The returned instance is a pointer to [None].
func NewNone[T any]() Optional[T] {
	return &None[T]{}
}

// IsPresent always returns false as [None] never contains a value.
func (n *None[T]) IsPresent() bool { return false }

// IsEmpty always returns true as [None] represents the absence of a value.
func (n *None[T]) IsEmpty() bool { return true }

// Get always returns the zero value of type T since [None] contains no value.
func (n *None[T]) Get() T { var zero T; return zero }
