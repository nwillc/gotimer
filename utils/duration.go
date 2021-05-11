package utils

import (
	"fmt"
	"strings"
	"time"
)

// Format a time.Duration into a string for the format `HH:MM.SS`.
func Format(work time.Duration, rest time.Duration, enableRest bool) (string, error) {
	var sb strings.Builder

	sb.WriteString(buildTime(work))

	// Rest time
	if enableRest {
		sb.WriteString(buildTime(rest))
	}
	return sb.String(), nil
}

func buildTime(d time.Duration) string {
	var sb strings.Builder

	if d >= time.Hour {
		sb.WriteString(fmt.Sprintf("%02d:", int(d.Hours())))
	}

	if d >= time.Minute {
		d %= time.Hour
		sb.WriteString(fmt.Sprintf("%02d.", int(d.Minutes())))
	}

	d %= time.Minute
	sb.WriteString(fmt.Sprintf("%02d", int(d.Seconds())))

	return sb.String()
}
