package box

import (
	"testing"

	"github.com/atendi9/capivara/assert"
)

func TestVoidAndNull(t *testing.T) {
	// Verifies if NULL is of type Void
	var _ Void = NULL
	
	// Since Void is struct{}, a map using it can be checked
	voidMap := map[string]Void{"exists": NULL}
	assert.LengthMap(t, 1, voidMap)
}

func TestSome(t *testing.T) {
	// Test standard NewSome with a string value
	optStr := NewSome("GoLang")
	
	assert.True(t, optStr.IsPresent())
	assert.False(t, optStr.IsEmpty())
	assert.Equal(t, "GoLang", optStr.Get())

	// Test NewSome with an integer value
	optInt := NewSome(42)
	
	assert.True(t, optInt.IsPresent())
	assert.False(t, optInt.IsEmpty())
	assert.Equal(t, 42, optInt.Get())

	// Test behavior of nil pointer for Some
	var nilSome *Some[int]
	var optNil Optional[int] = nilSome

	assert.False(t, optNil.IsPresent())
	assert.True(t, optNil.IsEmpty())
	assert.Equal(t, 0, optNil.Get())
}

func TestNone(t *testing.T) {
	// Test NewNone for a string type
	optNoneStr := NewNone[string]()

	assert.False(t, optNoneStr.IsPresent())
	assert.True(t, optNoneStr.IsEmpty())
	assert.Equal(t, "", optNoneStr.Get())

	// Test NewNone for an integer type
	optNoneInt := NewNone[int]()

	assert.False(t, optNoneInt.IsPresent())
	assert.True(t, optNoneInt.IsEmpty())
	assert.Equal(t, 0, optNoneInt.Get())
}

func TestOptionalInterfaceSlices(t *testing.T) {
	// Testing with slice length checks
	optionals := []Optional[string]{
		NewSome("first"),
		NewNone[string](),
		NewSome("third"),
	}

	assert.LengthSlice(t, 3, optionals)
	assert.True(t, optionals[0].IsPresent())
	assert.False(t, optionals[1].IsPresent())
	assert.Equal(t, "third", optionals[2].Get())
}