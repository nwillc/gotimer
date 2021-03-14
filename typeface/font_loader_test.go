package typeface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBitmaps(t *testing.T) {
	count := 12
	assert.NotNil(t, bitmaps)

	type args struct {
		font   string
		height int
	}
	tests := []struct {
		name string
		args args
	}{
		{"Six", args{"6", 6}},
		{"Seven", args{"7", 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runes, err := readBitmaps(bitmaps, "bitmaps/"+tt.args.font)
			assert.NoError(t, err)
			assert.NotNil(t, runes)
			assert.Equal(t, count, len(runes))
			for i := 0; i < 10; i++ {
				fr, ok := runes[rune('0'+i)]
				assert.True(t, ok)
				assert.Len(t, fr, tt.args.height, "rune %d", rune('0'+i))
			}
		})
	}
}

func TestToCharName(t *testing.T) {
	r, err := toCharName("a48.txt")
	assert.NoError(t, err)
	assert.Equal(t, '0', r)
}
