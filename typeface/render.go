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
