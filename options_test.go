package rehttp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOptions(t *testing.T) {
	op := NewOptions()
	assert.Equal(t, op.DialerKeepAlive, DefaultDialerKeepAlive,
		"Expected - NewOptions() to initialize with default values")

	got := op.Headers.Get("Connection")
	assert.Equal(t, got, "Keep-Alive", "Expected {Connection => Keep-Alive} got %s", got)
}
