package box

import (
	"errors"
	"testing"

	"github.com/atendi9/capivara/assert"
)

func TestResultSuccess(t *testing.T) {
	res := NewSuccess("data processing completed")

	assert.True(t, res.IsSuccess())
	assert.False(t, res.IsFailure())
	assert.Equal(t, "data processing completed", res.Value())
	assert.NoError(t, res.Error())
}

func TestResultFailure(t *testing.T) {
	expectedErr := errors.New("database connection lost")
	res := NewFailure[string](expectedErr)

	assert.False(t, res.IsSuccess())
	assert.True(t, res.IsFailure())
	assert.Equal(t, "", res.Value())
	assert.Error(t, res.Error())
	assert.Equal(t, expectedErr, res.Error())
}
