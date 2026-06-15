package box

// Ternary evaluates a boolean condition and returns one of two values.
func Ternary[T any](
	condition bool,
	resultTrue,
	resultFalse T,
) T {
	if condition {
		return resultTrue
	}
	return resultFalse
}

// LazyTernary evaluates a boolean condition and executes one of two functions, returning its result.
func LazyTernary[T any](condition bool, fnTrue, fnFalse func() T) T {
	if condition {
		return fnTrue()
	}
	return fnFalse()
}

// All returns true if all given boolean conditions evaluate to true.
func All(conditions ...bool) bool {
	for _, cond := range conditions {
		if !cond {
			return false
		}
	}
	return true
}

// When executes the provided function if the boolean condition evaluates to true.
func When(condition bool, fn func()) {
	if condition {
		fn()
	}
}

// Unless executes the provided function if the boolean condition evaluates to false.
func Unless(condition bool, fn func()) {
	if !condition {
		fn()
	}
}
