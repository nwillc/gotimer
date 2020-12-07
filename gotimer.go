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

type area struct {
	width  int
	height int
}

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

	var screen tcell.Screen
	if screen, err = tcell.NewScreen(); err != nil {
		panic(err)
	}

	if err := screen.Init(); err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second)
		display(duration, screen)
		duration = duration - time.Second
		if duration <= 0 {
			break
		}
	}

	time.Sleep(5 * time.Second)
	screen.Fini()
}

func display(duration time.Duration, screen tcell.Screen) {
	screen.Clear()
	str, _ := utils.Format(duration)
	y := 1
	for _, c := range str {
		w, err := typeface.RenderRune(screen, c, typeface.Medium, 0, y)
		if err != nil {
			panic(err)
		}
		y += w + 2
	}
	screen.Show()
}
