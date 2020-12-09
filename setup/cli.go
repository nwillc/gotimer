package setup

import (
	"flag"
)

// Flags for the command line interface.
var Flags struct {
	Version *bool
	Time    *string
	Color   *string
}

func init() {
	Flags.Version = flag.Bool("version", false, "Display version.")
	Flags.Time = flag.String("time", "25m", "The time for the timer")
	Flags.Color = flag.String("color", "orangered", "The display color")
}
