package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDString(t *testing.T) {
	str := FormatID(230)
	assert.Equal(t, str, "230")

	str = FormatID("hello")
	assert.Equal(t, str, "hello")

	str = FormatID(map[string]any{"a": "hello", "b": "bye"})
	assert.Equal(t, str, "hello|bye")
}
