package typeface

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"reflect"
	"testing"
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
