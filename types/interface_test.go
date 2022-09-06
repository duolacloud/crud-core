package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDString(t *testing.T) {
	str := fmt.Sprintf("%v", int(230))
	assert.Equal(t, str, "230")

	str = fmt.Sprintf("%v", int32(22312390))
	assert.Equal(t, str, "22312390")

	str = fmt.Sprintf("%v", int64(1923812340022312390))
	assert.Equal(t, str, "1923812340022312390")

	str = fmt.Sprintf("%v", "1923812340022312390")
	assert.Equal(t, str, "1923812340022312390")
}
