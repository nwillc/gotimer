package typeface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMediumRunes(t *testing.T) {
	var runes = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ':', '.'}
	for _, r := range runes {
		_, ok := Medium[r]
		assert.True(t, ok, "Missing rune %c", r)
	}
}
