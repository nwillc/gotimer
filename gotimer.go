package main

import (
	"flag"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/gen/version"
	"github.com/nwillc/gotimer/setup"
	"github.com/nwillc/gotimer/typeface"
	"github.com/nwillc/gotimer/utils"
	"os"
	"time"
)

// paused indicates timer is paused
var paused = false

func main() {
	flag.Parse()
	if *setup.Flags.Version {
		fmt.Println("Version:", version.Version)
		os.Exit(0)
	}

	duration, err := time.ParseDuration(*setup.Flags.Time)
	if err != nil {
		panic(err)
	}

	_, ok := typeface.AvailableFonts[*setup.Flags.FontName]
	if !ok {
		panic("Unknown font " + *setup.Flags.FontName)
	}

	color, ok := tcell.ColorNames[*setup.Flags.Color]
	if !ok {
		panic(fmt.Errorf("color %s not known", *setup.Flags.Color))
	}
	var s tcell.Screen
	if s, err = tcell.NewScreen(); err != nil {
		panic(err)
	}

	if err := s.Init(); err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			if paused {
				continue
			}
			display(duration, s, color, *setup.Flags.FontName)
			duration = duration - time.Second
			if duration < 0 {
				s.Beep()
				break
			}
		}
	}()

	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				s.Fini()
				os.Exit(0)
			} else if ev.Rune() == ' ' {
				paused = !paused
			}
		}
	}
}

func display(duration time.Duration, s tcell.Screen, color tcell.Color, fontName string) {
	font, ok := typeface.AvailableFonts[fontName]
	if !ok {
		panic("font not available")
	}
	s.Clear()
	str, err := utils.Format(duration)
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
