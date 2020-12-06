package setup

import "flag"

var Flags struct{
	Version *bool
	Time *string
}

func init()  {
	Flags.Version = flag.Bool("version", false, "Display version.")
	Flags.Time = flag.String("time", "25:00", "The time for the timer")
}
