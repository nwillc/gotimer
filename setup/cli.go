package setup

import "flag"

// Flags for the command line interface.
var Flags struct {
	Version *bool
	Time    *string
}

func init() {
	Flags.Version = flag.Bool("version", false, "Display version.")
	Flags.Time = flag.String("time", "25m", "The time for the timer")
}
