package setup

import (
	"flag"
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/typeface"
	"os"
)

type Values struct {
	Version   *bool
	Time      *string
	ColorName *string
	FontName  *string
}

// Flags for the command line interface.
var Flags *FlagSetWithValues

type FlagSetWithValues struct {
	*flag.FlagSet
	Values Values
}

func init() {
	Flags = newFlagSetWithValues(os.Args[0])
}

var defaultFont = "7"
var defaultColor = "orangered"

func newFlagSetWithValues(name string) *FlagSetWithValues {
	mfs := &FlagSetWithValues{
		FlagSet: flag.NewFlagSet(name, flag.ContinueOnError),
		Values:  Values{},
	}
	mfs.Values.Version = mfs.Bool("version", false, "Display version.")
	mfs.Values.Time = mfs.String("time", "25m", "The time for the timer")
	mfs.Values.ColorName = mfs.MemberString("color", &defaultColor, "set color",
		func(s string) bool { _, ok := tcell.ColorNames[s]; return ok },
		func(s string) { mfs.Values.ColorName = &s })
	mfs.Values.FontName = mfs.MemberString("font", &defaultFont, "set font",
		func(s string) bool { _, ok := typeface.AvailableFonts[s]; return ok },
		func(s string) { mfs.Values.FontName = &s})
	return mfs
}

func (m *FlagSetWithValues) ParseOsArgs() {
	err := m.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(2)
	}
}

func (m *FlagSetWithValues) MemberString(name string, value *string, usage string, memberFn func(string) bool, setterFn func(string)) *string {
	m.FlagSet.Func(name, fmt.Sprintf("%s (default \"%s\")", usage, *value), func(s string) error {
		if !memberFn(s) {
			return fmt.Errorf("unkown %s", name)
		}
		setterFn(s)
		return nil
	})
	if ok := memberFn(*value); !ok {
		panic(fmt.Sprintf("illegal default %s for %s", *value, name))
	}
	return value
}
