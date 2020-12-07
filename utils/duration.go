package utils

import (
	"fmt"
	"strings"
	"time"
)

// Format a time.Duration into a string for the format `HH:MM.SS`.
func Format(d time.Duration) (string, error) {
	var sb strings.Builder

	if d > time.Hour {
		sb.WriteString(fmt.Sprintf("%02d:", int(d.Hours())))
	}

	if d > time.Minute {
		d = d % time.Hour
		sb.WriteString(fmt.Sprintf("%02d.", int(d.Minutes())))
	}

	d = d % time.Minute
	sb.WriteString(fmt.Sprintf("%02d", int(d.Seconds())))
	return sb.String(), nil
}
