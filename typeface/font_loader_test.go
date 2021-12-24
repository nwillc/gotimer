/*
 *  Copyright (c) 2021,  nwillc@gmail.com
 *
 *  Permission to use, copy, modify, and/or distribute this software for any
 *  purpose with or without fee is hereby granted, provided that the above
 *  copyright notice and this permission notice appear in all copies.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 *  WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 *  MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 *  ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 *  WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 *  ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 *  OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package typeface

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			runes := readBitmaps(bitmaps, "bitmaps/"+tt.args.font)
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
	r := toCharName("a48.txt")
	assert.Equal(t, '0', r)
}
