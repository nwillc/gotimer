package setup

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/typeface"
)

type Values struct {
	Version   *bool
	Time      *string
	ColorName *string
	FontName  *string
}

var defaultFont = "7"
var defaultColor = "orangered"

func NewFlagSetWithValues(name string, values *Values) *CliFlagSet {
	fs := NewCliFlagSet(name)
	values.Version = fs.Bool("version", false, "Display version.")
	values.Time = fs.String("time", "25m", "The time for the timer")
	values.ColorName = fs.MemberString("color", &defaultColor, "set color",
		func(s string) bool { _, ok := tcell.ColorNames[s]; return ok },
		func(s string) { values.ColorName = &s })
	values.FontName = fs.MemberString("font", &defaultFont, "set font",
		func(s string) bool { _, ok := typeface.AvailableFonts[s]; return ok },
		func(s string) { values.FontName = &s})
	return fs
}
