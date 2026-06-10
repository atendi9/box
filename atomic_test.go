package box

import (
	"sync"
	"testing"

	"github.com/atendi9/capivara/assert"
)

func TestAtomic(t *testing.T) {
	t.Run("Load default zero value", func(t *testing.T) {
		var atom Atomic[int]
		assert.Equal(t, 0, atom.Load())
	})

	t.Run("Store and Load", func(t *testing.T) {
		var atom Atomic[int]

		atom.Store(42)
		assert.Equal(t, 42, atom.Load())

		atom.Store(100)
		assert.Equal(t, 100, atom.Load())
	})

	t.Run("Swap", func(t *testing.T) {
		var atom Atomic[string]

		old := atom.Swap("first")
		assert.Equal(t, "", old)
		assert.Equal(t, "first", atom.Load())

		old = atom.Swap("second")
		assert.Equal(t, "first", old)
		assert.Equal(t, "second", atom.Load())
	})

	t.Run("CompareAndSwap", func(t *testing.T) {
		var atom Atomic[int]
		atom.Store(10)

		swapped := atom.CompareAndSwap(5, 20)
		assert.False(t, swapped)
		assert.Equal(t, 10, atom.Load())

		swapped = atom.CompareAndSwap(10, 20)
		assert.True(t, swapped)
		assert.Equal(t, 20, atom.Load())
	})

	t.Run("Concurrent access safety", func(t *testing.T) {
		var atom Atomic[int]
		atom.Store(0)

		var wg sync.WaitGroup
		iterations := 100

		for range iterations {
			wg.Go(func() {
				for {
					current := atom.Load()
					if atom.CompareAndSwap(current, current+1) {
						break
					}
				}
			})
		}

		wg.Wait()

		assert.Equal(t, iterations, atom.Load())
	})
}
