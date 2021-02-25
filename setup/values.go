package setup

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nwillc/gotimer/eflag"
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
func NewEFlagSet(name string, values *Values) *eflag.EFSet {
	efs := eflag.NewEFlagSet(name)
	values.Version = efs.Bool("version", false, "Display version.")
	values.Time = efs.String("time", "25m", "The time for the timer")
	var colors []string
	for k := range tcell.ColorNames {
		colors = append(colors, k)
	}
	values.ColorName = efs.MemberString("color", &defaultColor, "set color", colors,
		func(s string) { values.ColorName = &s })
	values.FontName = efs.MemberString("font", &defaultFont, "set font", typeface.FontNames,
		func(s string) { values.FontName = &s })
	return efs
}
