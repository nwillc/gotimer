package typeface

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var blank = tcell.StyleDefault.
	Foreground(tcell.ColorBlack).
	Background(tcell.ColorBlack)

var filled = tcell.StyleDefault.
	Foreground(tcell.ColorOrangeRed).
	Background(tcell.ColorBlack)

// RenderRune renders the FontRune for a specified rune in a Font at a given location.
func RenderRune(s tcell.Screen, r rune, font Font, x int, y int) (int, error) {
	fontRune, ok := font[r]
	if !ok {
		return 0, fmt.Errorf("no font char for rune %c", r)
	}

	for xx, line := range fontRune {
		for yy, point := range line {
			style := blank
			if point {
				style = filled
			}
			s.SetCell(y+yy, x+xx, style, 0x2588)
		}
	}
	return len(fontRune[0]), nil
}
