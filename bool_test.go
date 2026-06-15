package box

import (
	"fmt"
	"testing"

	"github.com/atendi9/capivara/assert"
)

func TestTernary(t *testing.T) {
	t.Run("true condition returns first value", func(t *testing.T) {
		got := Ternary(true, "yes", "no")
		assert.Equal(t, "yes", got)
	})

	t.Run("false condition returns second value", func(t *testing.T) {
		got := Ternary(false, "yes", "no")
		assert.Equal(t, "no", got)
	})

	t.Run("works with integers", func(t *testing.T) {
		got1 := Ternary(true, 1, 2)
		assert.Equal(t, 1, got1)

		got2 := Ternary(false, 1, 2)
		assert.Equal(t, 2, got2)
	})

	t.Run("works with boolean values", func(t *testing.T) {
		got1 := Ternary(true, true, false)
		assert.True(t, got1)

		got2 := Ternary(false, true, false)
		assert.False(t, got2)
	})

	t.Run("works with maps", func(t *testing.T) {
		mapA := map[string]any{"age": 25}
		mapB := map[string]any{"age": 30, "height": 180}

		got := Ternary(true, mapA, mapB)
		assert.Equal(t, fmt.Sprint(mapA), fmt.Sprint(got))
		assert.LengthMap(t, 1, got)
	})

	t.Run("works with slices", func(t *testing.T) {
		sliceA := []string{"Gabriel Luiz", "Gediael"}
		sliceB := []string{"John"}

		got := Ternary(true, sliceA, sliceB)
		assert.Equal(t, fmt.Sprint(sliceA), fmt.Sprint(got))
		assert.LengthSlice(t, 2, got)
	})

	t.Run("works with custom structs", func(t *testing.T) {
		type user struct {
			name string
			role string
		}
		u1 := user{name: "Alice", role: "admin"}
		u2 := user{name: "Bob", role: "user"}

		got := Ternary(false, u1, u2)
		assert.Equal(t, u2, got)
	})

	t.Run("works with pointers", func(t *testing.T) {
		val1 := 42
		val2 := 99

		got := Ternary(true, &val1, &val2)
		assert.Equal(t, &val1, got)
	})

	t.Run("works with float values", func(t *testing.T) {
		got := Ternary(false, 3.14, 9.81)
		assert.Equal(t, 9.81, got)
	})
}

func TestLazyTernary(t *testing.T) {
	t.Run("true condition returns result of first function", func(t *testing.T) {
		got := LazyTernary(true, func() string { return "yes" }, func() string { return "no" })
		assert.Equal(t, "yes", got)
	})

	t.Run("false condition returns result of second function", func(t *testing.T) {
		got := LazyTernary(false, func() int { return 1 }, func() int { return 2 })
		assert.Equal(t, 2, got)
	})

	t.Run("evaluates lazily and prevents execution of unselected function", func(t *testing.T) {
		trueCalled := false
		falseCalled := false

		fnTrue := func() int {
			trueCalled = true
			return 1
		}

		fnFalse := func() int {
			falseCalled = true
			return 2
		}

		got := LazyTernary(true, fnTrue, fnFalse)

		assert.Equal(t, 1, got)
		assert.True(t, trueCalled)
		assert.False(t, falseCalled)
	})
}

func TestAll(t *testing.T) {
	t.Run("all true conditions return true", func(t *testing.T) {
		got := All(true, true, true)
		assert.True(t, got)
	})

	t.Run("one false condition returns false", func(t *testing.T) {
		got := All(true, false, true)
		assert.False(t, got)
	})

	t.Run("all false conditions return false", func(t *testing.T) {
		got := All(false, false)
		assert.False(t, got)
	})

	t.Run("empty conditions return true", func(t *testing.T) {
		got := All()
		assert.True(t, got)
	})
}

func TestWhen(t *testing.T) {
	t.Run("true condition executes function", func(t *testing.T) {
		called := false
		When(true, func() {
			called = true
		})
		assert.True(t, called)
	})

	t.Run("false condition does not execute function", func(t *testing.T) {
		called := false
		When(false, func() {
			called = true
		})
		assert.False(t, called)
	})
}

func TestUnless(t *testing.T) {
	t.Run("false condition executes function", func(t *testing.T) {
		called := false
		Unless(false, func() {
			called = true
		})
		assert.True(t, called)
	})

	t.Run("true condition does not execute function", func(t *testing.T) {
		called := false
		Unless(true, func() {
			called = true
		})
		assert.False(t, called)
	})
}
