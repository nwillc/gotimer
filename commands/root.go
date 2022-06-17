/*
 *  Copyright (c) 2022,  nwillc@gmail.com
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

package commands

import (
	"fmt"
	"github.com/nwillc/genfuncs/container/gslices"
	"os"
	"time"

	"github.com/nwillc/genfuncs"
	"github.com/nwillc/genfuncs/container"
	"github.com/nwillc/gotimer/generics"

	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/display"
	"github.com/nwillc/gotimer/gen/version"
	"github.com/nwillc/gotimer/typeface"
	"github.com/spf13/cobra"
)

var (
	cliValues struct {
		Version   bool
		Time      string
		ColorName string
		FontName  string
	}
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&cliValues.Version, "version", "v", false, "Display version")
	RootCmd.PersistentFlags().StringVarP(&cliValues.Time, "time", "t", "25m", "Time to count down")
	RootCmd.PersistentFlags().StringVarP(&cliValues.ColorName, "color", "c", "orangered", "Color of timer")
	RootCmd.PersistentFlags().StringVarP(&cliValues.FontName, "size", "s", "7", "Font size to use")
}

// RootCmd is the root, and only, command of gotimer.
var RootCmd = &cobra.Command{
	Use:   "gotimer",
	Short: "A digital count down timer",
	Long:  "A simple terminal based digital count down timer, may be used as a Pomodoro timer.",
	Args:  cobra.ExactArgs(0),
	Run:   timerCmd,
}

func timerCmd(_ *cobra.Command, _ []string) {
	if cliValues.Version {
		fmt.Println("Version:", version.Version)
		os.Exit(0)
	}
	duration, err := time.ParseDuration(cliValues.Time)
	if err != nil {
		fmt.Println("Illegal time:", err)
		fmt.Println("A time ia a sequence of decimal numbers, each with optional fraction and a unit suffix,")
		fmt.Println(`such as ".5h", "1m35s" or "2h45m". Valid time units are "s", "m", "h".`)
		os.Exit(1)
	}
	color, ok := tcell.ColorNames[cliValues.ColorName]
	if !ok {
		fmt.Println("Unknown color:", cliValues.ColorName)
		fmt.Println("Available colors:", colors(tcell.ColorNames))
		os.Exit(1)
	}
	font, ok := typeface.AvailableFonts[cliValues.FontName]
	if !ok {
		fmt.Println("Unknown font:", cliValues.FontName)
		fmt.Println("Available font sizes:", fonts(typeface.FontNames))
		os.Exit(1)
	}
	display.Timer(duration, color, font)
}

func colors(colorNames container.GMap[string, tcell.Color]) string {
	return colorNames.
		Keys().
		Filter(genfuncs.IsNotBlank).
		SortBy(genfuncs.OrderedLess[string]).
		JoinToString(generics.AToA, "\n  ", "\n  ", "")
}

func fonts(names container.GSlice[string]) string {
	return gslices.Map(names, generics.AToi).
		SortBy(genfuncs.OrderedLess[int]).
		JoinToString(generics.IToA, "\n  ", "\n  ", "")
}
