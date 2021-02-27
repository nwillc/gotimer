package setup

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/efs"
	"github.com/nwillc/gotimer/typeface"
)

// Values settable from the command line.
type Values struct {
	Version   *bool
	Time      *string
	ColorName *string
	FontName  *string
}

var defaultFont = "7"
var defaultColor = "orangered"

// NewEFlagSet sets up a EFlagSet for the Values provided.
func NewEFlagSet(name string, values *Values) *efs.EFlagSet {
	fs := efs.NewEFlagSet(name)
	values.Version = fs.Bool("version", false, "Display version.")
	values.Time = fs.String("time", "25m", "The time for the timer")
	var colors []string
	for k := range tcell.ColorNames {
		colors = append(colors, k)
	}
	values.ColorName = fs.MemberString("color", &defaultColor, "set color", colors,
		func(s string) { values.ColorName = &s })
	values.FontName = fs.MemberString("font", &defaultFont, "set font", typeface.FontNames,
		func(s string) { values.FontName = &s })
	return fs
}
