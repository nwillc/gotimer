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

package display

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/typeface"
)

// Timer runs for given time.Duration, tcell.Color, and typeface.Font.
func Timer(duration time.Duration, color tcell.Color, font typeface.Font) {
	// Prepare the display
	var s tcell.Screen
	var err error
	if s, err = tcell.NewScreen(); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}

	// paused indicates timer is paused
	var paused = false
	// Display the timer
	stop := RepeatUntilStopped(time.Second, func() {
		if paused || duration < 0 {
			return
		}
		display(duration, s, color, font)
		duration -= time.Second
		if duration < 0 {
			_ = s.Beep()
		}
	})
	// Process keyboard
	for {
		ev := s.PollEvent()
		switch et := ev.(type) {
		case *tcell.EventKey:
			if et.Key() == tcell.KeyEscape || et.Key() == tcell.KeyCtrlC {
				stop <- true
				s.Fini()
				os.Exit(0)
			} else if et.Rune() == ' ' {
				paused = !paused
			}
		}
	}
}

func display(duration time.Duration, s tcell.Screen, color tcell.Color, font typeface.Font) {
	s.Clear()
	str, err := Format(duration)
	if err != nil {
		panic(err)
	}
	x := 1
	for _, c := range str {
		width, err := typeface.RenderRune(s, c, font, color, x, 1)
		if err != nil {
			panic(err)
		}
		x += width + 1
	}
	s.Show()
}
