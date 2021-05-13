[![license](https://img.shields.io/github/license/nwillc/gotimer.svg)](https://tldrlegal.com/license/-isc-license)
[![CI](https://github.com/nwillc/gotimer/workflows/CI/badge.svg)](https://github.com/nwillc/gotimer/actions?query=workflow%3CI)
[![Go Report Card](https://goreportcard.com/badge/github.com/nwillc/gotimer)](https://goreportcard.com/report/github.com/nwillc/gotimer)
---
# Go Timer

A simple terminal based digital timer written in Go to use as a Pomodoro time.

![gotimer](gotimer.png)

## With resting option

![gotimer with rest](gotimer_rest.png)

## Usage:

```bash
$ ./gotimer -h
Usage of ./gotimer:
  -color value
    	set color (default "orangered")
  -colorRest value
    	set color (default "skyblue")
  -font value
    	set font (default "7")
  -time string
    	The time for the timer (default "25m")
  -timeRest string
    	The time for the rest timer (default "5m")
  -version
    	Display version.
```

 - The time duration is given in hours, minutes and seconds: `#h#m#s`.
 - Exit with Ctrl-C or ESC
 - Pause with SPACE
