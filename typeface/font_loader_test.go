package typeface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitmaps(t *testing.T) {
	count := 12
	assert.NotNil(t, bitmaps)
	runes, err := readBitmaps(bitmaps, "bitmaps/medium")
	assert.NoError(t, err)
	assert.NotNil(t, runes)
	assert.Equal(t, count, len(runes))
	for i := 0; i < 10; i++ {
		fr, ok := runes[rune('0' + i)]
		assert.True(t, ok)
		assert.Len(t, fr, 6)
	}
}

func TestToCharName(t *testing.T) {
	r, err := toCharName("a48.txt")
	assert.NoError(t, err)
	assert.Equal(t, '0', r)
}
