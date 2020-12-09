package typeface

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

var blockOff = style(tcell.ColorBlack, tcell.ColorBlack)

// RenderRune renders the FontRune for a specified rune in a Font at a given location.
func RenderRune(s tcell.Screen, r rune, font Font, c tcell.Color, x int, y int) (int, error) {
	fontRune, ok := font[r]
	if !ok {
		return 0, fmt.Errorf("no font char for rune %c", r)
	}
	blockOn := style(c, tcell.ColorBlack)
	for yy, line := range fontRune {
		for xx, point := range line {
			style := blockOff
			if point {
				style = blockOn
			}
			s.SetCell(x+xx, y+yy, style, 0x2588)
		}
	}
	return len(fontRune), nil
}

func style(foreground tcell.Color, background tcell.Color) tcell.Style {
	return tcell.StyleDefault.Foreground(foreground).Background(background)
}
