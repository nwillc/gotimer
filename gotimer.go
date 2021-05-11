package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/gen/version"
	"github.com/nwillc/gotimer/setup"
	"github.com/nwillc/gotimer/typeface"
	"github.com/nwillc/gotimer/utils"
)

func main() {
	flags := &setup.Values{}
	setup.NewEFlagSet(os.Args[0], flags).ParseOsArgs()
	if *flags.Version {
		fmt.Println("Version:", version.Version)
		os.Exit(0)
	}

	duration, err := time.ParseDuration(*flags.Time)
	if err != nil {
		panic(err)
	}
	durationRest, err := time.ParseDuration(*flags.TimeRest)
	if err != nil {
		panic(err)
	}

	color := tcell.ColorNames[*flags.ColorName]
	colorRest := tcell.ColorNames[*flags.ColorNameRest]
	var s tcell.Screen
	if s, err = tcell.NewScreen(); err != nil {
		panic(err)
	}

	if err := s.Init(); err != nil {
		panic(err)
	}

	// paused indicates timer is paused
	var paused = false
	go func() {
		for {
			time.Sleep(time.Second)
			if paused {
				continue
			}
			display(duration, durationRest, s, color, colorRest, *flags.FontName)
			duration -= time.Second
			if duration < 0 {
				_ = s.Beep()
				break
			}
		}
	}()

	for {
		ev := s.PollEvent()
		switch et := ev.(type) {
		case *tcell.EventKey:
			if et.Key() == tcell.KeyEscape || et.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			} else if et.Rune() == ' ' {
				paused = !paused
			}
		}
	}
}

func display(duration time.Duration, durationRest time.Duration, s tcell.Screen, color tcell.Color, colorRest tcell.Color, fontName string) {
	var width int
	font, ok := typeface.AvailableFonts[fontName]
	if !ok {
		panic("font not available")
	}
	s.Clear()
	str, err := utils.Format(duration, durationRest, true)
	if err != nil {
		panic(err)
	}
	x := 1
	state := "work"
	for _, c := range str {
		if c == ' ' {
			state = "rest"
		}
		if state == "work" {
			width, err = typeface.RenderRune(s, c, font, color, x, 1)
		} else {
			width, err = typeface.RenderRune(s, c, font, colorRest, x, 1)
		}
		if err != nil {
			panic(err)
		}
		x += width + 1
	}
	s.Show()
}
