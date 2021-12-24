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
	_ "embed"
	"reflect"
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestRenderRune(t *testing.T) {
	type args struct {
		s    tcell.Screen
		r    rune
		font Font
		c    tcell.Color
		x    int
		y    int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RenderRune(tt.args.s, tt.args.r, tt.args.font, tt.args.c, tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("RenderRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RenderRune() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_style(t *testing.T) {
	type args struct {
		foreground tcell.Color
		background tcell.Color
	}
	tests := []struct {
		name string
		args args
		want tcell.Style
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := style(tt.args.foreground, tt.args.background); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("style() = %v, want %v", got, tt.want)
			}
		})
	}
}
